package main

import (
	"encoding/json"
	"fmt"
	"golang-kafka-client/conf"
	"golang-kafka-client/db"
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
	conn     *db.DB
}

func (kafkaClient *KafkaClient) Init(config conf.Config, db *db.DB) {
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
}

func (kafkaClient *KafkaClient) logConsume(str []byte) {
	msg := db.Tb_co_log{}
	json.Unmarshal(str, &msg)

	// DB INSERT
	tbCoLogDao := db.TbCoLogDao{Conn: kafkaClient.conn}

	tbCoLogDao.Create(msg)
}
