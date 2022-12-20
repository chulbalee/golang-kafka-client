package main

import (
	"fmt"
	"golang-kafka-client/types"
)

func DBInit(config types.Config) {
	fmt.Println(":::db Init")
	fmt.Println(":::db type ", config.DbConfig.Type)
	fmt.Println(":::db host ", config.DbConfig.Host)
}
