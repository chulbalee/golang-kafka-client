package main

import (
	"encoding/json"
	"fmt"
	"golang-kafka-client/types"
	"os"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	MIN_COMMIT_COUNT = 1
)

type KafkaClient struct {
	BootstrapServers string
	GroupId          string
	Topics           []string

	consumer *kafka.Consumer
}

type logMsg struct {
	Tx  string `json:"tx"`
	Id  string `json:"id"`
	Msg string `json:"msg"`
}

func (kafkaClient *KafkaClient) Init(config types.Config) *KafkaClient {
	fmt.Println(":::kafka Init")
	fmt.Println(":::kafka bootstrap-servers ", strings.Join(config.KafkaClientConfig.BootstrapServers, ","))
	fmt.Println(":::kafka GroupId ", config.KafkaClientConfig.GroupId)

	kafkaClient.BootstrapServers = strings.Join(config.KafkaClientConfig.BootstrapServers, ",")
	kafkaClient.GroupId = config.KafkaClientConfig.GroupId
	kafkaClient.Topics = append(kafkaClient.Topics, config.KafkaClientConfig.Topics)

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        kafkaClient.BootstrapServers,
		"broker.address.family":    "v4",
		"group.id":                 kafkaClient.GroupId,
		"session.timeout.ms":       6000,
		"auto.offset.reset":        "earliest",
		"enable.auto.offset.store": false,
	})

	if err != nil {
		panic(err)
	}

	kafkaClient.consumer = consumer

	return kafkaClient
}

func (kafkaClient *KafkaClient) Run() {
	_ = kafkaClient.consumer.SubscribeTopics(kafkaClient.Topics, nil)

	run := true
	msg_count := 0
	for run {
		ev := kafkaClient.consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			msg_count = (msg_count + 1) % 100
			if msg_count%MIN_COMMIT_COUNT == 0 {
				kafkaClient.consumer.Commit()
			}

			switch *e.TopicPartition.Topic {
			case "log":
				go logConsume(e.Value)

			default:
				fmt.Printf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value))
			}
		case kafka.PartitionEOF:
			fmt.Printf("%% Reached %v\n", e)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			fmt.Printf("Ignored %v\n", e)
		}
	}
}

func logConsume(str []byte) {
	msg := logMsg{}
	json.Unmarshal(str, &msg)

	// DB INSERT
}
