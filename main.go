package main

import (
	"golang-kafka-client/conf"
	"golang-kafka-client/db"
)

const (
	CONFIG_PATH_PREFIX = "./conf/"
	CONFIG_EXTENSION   = ".yaml"
	SERVER_CONFIG_PATH = CONFIG_PATH_PREFIX + "server" + CONFIG_EXTENSION
)

func main() {
	var config conf.Config

	config.LoadConfig(SERVER_CONFIG_PATH)

	db.Init(config)

	kafkaClient := KafkaClient{}

	kafkaClient.Init(config, db.GetInstance())
	kafkaClient.Run()
}
