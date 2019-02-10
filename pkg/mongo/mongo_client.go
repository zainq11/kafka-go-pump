package mongo

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"time"
)

type Content struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Data        string `json:"content"`
	CreatedTime string `json:"time"`
}

type Config struct {
	DbUrl  string
	DbName string
}

type Client struct {
	client   *mongo.Client
	database *mongo.Database
}

func CreateClient(cfg Config) (*Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, cfg.DbUrl)
	return &Client{client, client.Database(cfg.DbName)}, err
}

func (c *Client) SaveContent(content Content, collectionName string) error {
	collection := c.database.Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if res, err := collection.InsertOne(ctx, bson.M{
		"title":    content.Title,
		"data":     content.Data,
		"author":   content.Author,
		"time":     content.CreatedTime}); err != nil {
		fmt.Errorf("An error occured when trying to save document")
		return err
	} else {
		fmt.Printf("Inserted document successfully with id: %d", res.InsertedID)
		return nil
	}

}
