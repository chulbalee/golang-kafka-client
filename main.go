package main

import (
	"golang-kafka-client/types"
)

const (
	CONFIG_PATH_PREFIX = "./conf/"
	CONFIG_EXTENSION   = ".yaml"
	SERVER_CONFIG_PATH = CONFIG_PATH_PREFIX + "server" + CONFIG_EXTENSION
)

func main() {
	var config types.Config
	config.LoadConfig(SERVER_CONFIG_PATH)

	DBInit(config)

	kafkaClient := KafkaClient{}

	kafkaClient.Init(config)
	kafkaClient.Run()
}
