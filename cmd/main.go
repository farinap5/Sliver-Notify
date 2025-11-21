package main

import (
	"connect/internal"
	"log"
)

func main() {
	cfg, err := internal.NewConfig("./config.yaml")
	if err != nil {
		log.Println(err.Error())
		return
	}
	cfg.Start()
}