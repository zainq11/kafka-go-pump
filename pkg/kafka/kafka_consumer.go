package kafka

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"strings"
	"time"
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
	Reader   *kafka.Reader
}

type doOnRead func(value []byte)

func CreateConsumer(cfg Config) (Consumer, error) {
	// Make a new reader that consumes from topic-A
	config := kafka.ReaderConfig{
		Brokers:         strings.Split(cfg.KafkaBrokerUrl, ","),
		GroupID:         cfg.KafkaClientId,
		Topic:           cfg.KafkaTopic,
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: -1,
	}

	return Consumer{KafkaCfg: cfg, Reader: kafka.NewReader(config)}, nil
}

func (consumer *Consumer) Subscribe(callback doOnRead) {
	reader := consumer.Reader
	// Close Reader once this program exits
	defer reader.Close()

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Error().Msgf("Error while receiving message: %s", err.Error())
			continue
		}
		value := m.Value
		log.Printf("message at topic/partition/offset %v/%v/%v: %s\n", m.Topic, m.Partition, m.Offset, string(value))
		callback(value)
	}
}
