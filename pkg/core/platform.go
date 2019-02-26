package core

import (
	"encoding/json"
	"fmt"
	"github.com/zianKazi/social-content-data-service/pkg/kafka"
	"github.com/zianKazi/social-content-data-service/pkg/mongo"
)

type Properties struct {
	BrokerUrl string
	Client    *mongo.Client
}

// Represent a platform
type PlatformContext struct {
	Name           string
	TopicName      string
	CollectionName string
	client         *mongo.Client
	parser         func([]byte) map[string]interface{}
	//consumer       *kafka.Consumer
}

type PlatformMap map[string]PlatformContext

const (
	TWITTER string = "twitter"
	REDDIT  string = "reddit"
)

//type PlatformHandle interface {
//	Run(props Properties) error
//}

func (context *PlatformContext) Boot(props Properties) error {
	var config = kafka.Config{KafkaTopic: context.TopicName, KafkaConsumerGroup: context.Name, KafkaClientId: context.Name + "xx", KafkaBrokerUrl: props.BrokerUrl}
	if consumer, error := kafka.CreateConsumer(config); error != nil {
		fmt.Println("Failed to initialize the Kafka Consumer")
		return error
	} else {
		//context.consumer = &consumer
		consumer.Subscribe(func(value []byte) {
			//document := mongo.Document{Collection: context.CollectionName, }
			//json.Unmarshal(value, &content)
			go context.client.SaveContent(context.CollectionName, context.parser(value))
		})
		return nil
	}
}

func CreatePlatformMap(props Properties) (PlatformMap, error) {
	platformMap := PlatformMap{}
	reddit := PlatformContext{
		Name: REDDIT, TopicName: REDDIT, CollectionName: "reddit",
		client: props.Client,
		parser: func(bytes []byte) map[string]interface{} {
			var content map[string]interface{}
			json.Unmarshal(bytes, &content)
			return content
		}}
	if error := reddit.Boot(props); error != nil {
		fmt.Println("Failed to initialize the PlatformContext for " + reddit.Name)
		return nil, error
	} else {
		platformMap[reddit.Name] = reddit
		return platformMap, nil
	}
}
