package internal

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/bishopfox/sliver/client/assets"
	"github.com/bishopfox/sliver/client/constants"
	"github.com/bishopfox/sliver/client/transport"
	"github.com/bishopfox/sliver/protobuf/clientpb"
	"github.com/bishopfox/sliver/protobuf/commonpb"
	"google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v2"
)

type Chat struct {
	Token string `yaml:"token"`
	Chat  string `yaml:"chat"`
}

type Config struct {
	Path     	string 	`yaml:"path"`
	Telegram 	Chat 	`yaml:"telegram"`
	Discord 	Chat 	`yaml:"discord"`
	Teams struct {
		Webhook string 	`yaml:"webhook"`
	} `yaml:"teams"`
	Message 	string	`yaml:"-"`
}

func NewConfig(path string) (Config, error) {
	cfg := Config{}

	b, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, err
}

func (c Config)Start() {
	log.Println("Starting Sliver Connect")
	conf, err := assets.ReadConfig(c.Path)
	if err != nil {
		log.Printf("%s\n", err.Error())
		return
	}

	log.Printf("Connect to server %s on port %d\n", conf.LHost, conf.LPort)
	rpc, ln, err := transport.MTLSConnect(conf)
	if err != nil {
		log.Printf("%s\n", err.Error())
		return
	}
	defer ln.Close()

	eventStream, err := rpc.Events(context.Background(), &commonpb.Empty{})
	if err != nil {
		log.Printf("%s\n", err.Error())
		return
	}


	for {
		event, err := eventStream.Recv()
		if err == io.EOF || event == nil {
			log.Printf("%s\n", err.Error())
			return
		}
		msg := `
UUID: %s 
BINARY: %s
USER: %s
MACHINE: %s (%s)
ARCH: %s/%s`

		switch event.EventType {
			case constants.SessionOpenedEvent:
				session := event.Session
				shortID := strings.Split(session.ID, "-")[0]

				msg = "New Session " + msg
				d := fmt.Sprintf(msg, shortID, session.Name, session.Username, session.RemoteAddress, session.Hostname, session.OS, session.Arch)

				log.Println(d)
				c.SendData(d)

			case constants.BeaconRegisteredEvent:
				beacon := &clientpb.Beacon{}
				proto.Unmarshal(event.Data, beacon)
				shortID := strings.Split(beacon.ID, "-")[0]

				msg = "New Beacon " + msg
				d := fmt.Sprintf(msg, shortID, beacon.Name, beacon.Username, beacon.RemoteAddress, beacon.Hostname, beacon.OS, beacon.Arch)
				
				log.Println(d)
				c.SendData(d)
		}
	}
}