package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/VladKuch/go_practic/internal/app/apiserver"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error init env variables: %s", err.Error())
	}

	flag.Parse()

	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatalf("Error server init: %s", err.Error())
	}
}
