package main

import (
	"connect/internal"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		println("Use: ./sliver-notify config.yaml")
		return
	}
	 
	cfg, err := internal.NewConfig(os.Args[1])
	if err != nil {
		log.Println(err.Error())
		return
	}
	cfg.Start()
}