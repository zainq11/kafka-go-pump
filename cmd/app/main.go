package main

import (
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
	//TODO: Use arguments
	flag.StringVar(&mongoCfg.DbUrl, "db-url", "mongodb://127.0.0.1:27017", "Full connection url for database")
	flag.StringVar(&mongoCfg.DbName, "db-name", "trendDB", "Database Name")

	flag.StringVar(&kafkaCfg.KafkaBrokerUrl, "kafka-brokers", "localhost:19092,localhost:29092,localhost:39092", "Kafka brokers in comma separated value")
	flag.BoolVar(&kafkaCfg.KafkaVerbose, "kafka-verbose", true, "Kafka verbose logging")
	flag.StringVar(&kafkaCfg.KafkaTopic, "kafka-topic", "rdt", "Kafka topic. Only one topic per worker.")
	flag.StringVar(&kafkaCfg.KafkaConsumerGroup, "kafka-consumer-group", "rdt-data", "Kafka consumer group")
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
		if platformMap, error := core.CreatePlatformMap(core.Properties{kafkaCfg.KafkaBrokerUrl, client}); error != nil {
			fmt.Println("Failed to initialize the Platform map")
		}else {
			fmt.Println("Initialized Platform map " + string(len(platformMap)))
		}
	}

}
