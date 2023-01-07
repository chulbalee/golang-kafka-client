package main

import (
	"encoding/json"
	"fmt"
	"golang-kafka-client/conf"
	"golang-kafka-client/db"
	"os"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gorm.io/gorm"
)

const (
	MIN_COMMIT_COUNT = 1
)

type KafkaClient struct {
	BootstrapServers string
	GroupId          string
	Topics           []string

	consumer *kafka.Consumer
	conn     *gorm.DB
}

func (kafkaClient *KafkaClient) Init(config conf.Config, db *gorm.DB) {
	fmt.Println(":::kafka Init")
	fmt.Println(":::kafka bootstrap-servers ", strings.Join(config.KafkaClient.BootstrapServers, ","))
	fmt.Println(":::kafka GroupId ", config.KafkaClient.GroupId)

	kafkaClient.BootstrapServers = strings.Join(config.KafkaClient.BootstrapServers, ",")
	kafkaClient.GroupId = config.KafkaClient.GroupId
	kafkaClient.Topics = config.KafkaClient.Topics

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
	kafkaClient.conn = db
}

func (kafkaClient *KafkaClient) Run() {
	fmt.Println("::: Kafka Topic Consume => ", kafkaClient.Topics)
	err := kafkaClient.consumer.SubscribeTopics(kafkaClient.Topics, nil)

	if err != nil {
		panic(err)
	}

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
			case "TB_CO_LOG_HIST":
				go kafkaClient.logConsume(e.Value)

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

	kafkaClient.consumer.Close()
}

func (kafkaClient *KafkaClient) logConsume(str []byte) {
	fmt.Println("::: LOG CONSUMED")
	var msg string
	json.Unmarshal(str, &msg)

	entity := db.Tb_co_log{}
	entity.BasDt = time.Now().Format("20060102")
	entity.Msg = msg

	fmt.Println("::: LOG CONSUMED : ", entity)
	// DB INSERT
	db.Create(entity)
}
