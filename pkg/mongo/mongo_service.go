package mongo

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/zianKazi/social-content-data-service/pkg/core"
	"time"
)

type Client struct {
	client   *mongo.Client
	database *mongo.Database
}

func CreateClient(url string, dbName string) (*Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, url)
	return &Client{client, client.Database(dbName)}, err
}

func (c *Client) SaveContent(content core.Content, collectionName string) error {
	collection := c.database.Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if res, err := collection.InsertOne(ctx, bson.M{
		"title":    content.Title,
		"data":     content.Data,
		"author":   content.Author,
		"platform": content.Platform,
		"time":     content.CreatedDate}); err != nil {
		fmt.Errorf("An error occured when trying to save document")
		return err
	} else {
		fmt.Printf("Inserted document successfully with id: %d", res.InsertedID)
		return nil
	}

}
