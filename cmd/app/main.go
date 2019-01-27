package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/zianKazi/social-content-data-service/pkg/core"
	"github.com/zianKazi/social-content-data-service/pkg/kafka"
	"github.com/zianKazi/social-content-data-service/pkg/mongo"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Init config holders
	var kafkaCfg = kafka.Config{}
	var mongoCfg = mongo.Config{}

	// Parse flag values
	flag.StringVar(&mongoCfg.DbUrl, "db-url", "", "Full connection url for database")
	flag.StringVar(&mongoCfg.DbName, "db-name", "", "Database Name")

	flag.StringVar(&kafkaCfg.KafkaBrokerUrl, "kafka-brokers", "localhost:19092,localhost:29092,localhost:39092", "Kafka brokers in comma separated value")
	flag.BoolVar(&kafkaCfg.KafkaVerbose, "kafka-verbose", true, "Kafka verbose logging")
	flag.StringVar(&kafkaCfg.KafkaTopic, "kafka-topic", "foo", "Kafka topic. Only one topic per worker.")
	flag.StringVar(&kafkaCfg.KafkaConsumerGroup, "kafka-consumer-group", "consumer-group", "Kafka consumer group")
	flag.StringVar(&kafkaCfg.KafkaClientId, "kafka-client-id", "my-client-id", "Kafka client id")

	flag.Parse()

	//TODO: echo all config values
	fmt.Println("Flag %s is %s", "url", mongoCfg.DbUrl)
	fmt.Println("Flag %s is %s", "db", mongoCfg.DbName)
	fmt.Println("Flag %s is %s", "broker", kafkaCfg.KafkaBrokerUrl)

	//TODO: evaluate this
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	if client, error := mongo.CreateClient(mongoCfg); error != nil {
		fmt.Println("Failed to initialize the Mongo Client")
		return
	} else {

		if consumer, error := kafka.CreateConsumer(kafkaCfg); error != nil {
			fmt.Println("Failed to initialize the Kafka Consumer")
			return
		} else {
			consumer.Subscribe(func(value []byte) {
				content := core.Content{}
				json.Unmarshal(value, &content)
				go client.SaveContent(content, "reddit-content")
			})
		}
	}

	//} else {
	//	client.SaveContent(core.Content{
	//		Title:       "zain's first",
	//		Author:      "Zain Qazi",
	//		CreatedDate: time.Now(),
	//		Data:        "Here we are!",
	//		Platform:    "Reddit"}, "testCollection")
	//}
}
