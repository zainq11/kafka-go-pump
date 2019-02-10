package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	//"strings"
)

type Config struct {
	KafkaBrokerUrl     string
	KafkaVerbose       bool
	KafkaTopic         string
	KafkaConsumerGroup string
	KafkaClientId      string
}

type Consumer struct {
	KafkaCfg Config
	Reader   *kafka.Consumer
}

type doOnRead func(value []byte)

func CreateConsumer(cfg Config) (Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.KafkaBrokerUrl,
		"group.id":          cfg.KafkaConsumerGroup,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{cfg.KafkaTopic}, nil)

	return Consumer{KafkaCfg: cfg, Reader: c}, nil
}

func (consumer *Consumer) Subscribe(callback doOnRead) {
	for {
		msg, err := consumer.Reader.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			callback(msg.Value)
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	consumer.Reader.Close()
}
