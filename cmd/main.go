package main

import (
	"flag"
	"img-generator/configs"
	"img-generator/internal/server"
	"log"
)

var confPath = flag.String("conf-path", "./configs/.env", "Path to config env.")

func main() {
	conf, err := configs.New(*confPath)
	if err != nil {
		log.Fatalln(err)
	}
	server.Run(conf)
}
