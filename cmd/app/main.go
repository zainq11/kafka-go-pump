package main

import (
	"flag"
	"fmt"
	"github.com/zianKazi/social-content-data-service/pkg/core"
	"github.com/zianKazi/social-content-data-service/pkg/mongo"
	"time"
)

func main() {
	var url = flag.String("url", "", "Connection Url for database")
	var db = flag.String("db", "", "Database name")

	flag.Parse()

	fmt.Println("Flag %s is %s", "url", *url)
	fmt.Println("Flag %s is %s", "db", *db)

	if client, error := mongo.CreateClient(*url, *db); error != nil {
		fmt.Println("Failed to initialize the client")
		return
	} else {
		client.SaveContent(core.Content{
			Title:       "zain's first",
			Author:      "Zain Qazi",
			CreatedDate: time.Now(),
			Data:        "Here we are!",
			Platform:    "Reddit"}, "testCollection")
	}
}
