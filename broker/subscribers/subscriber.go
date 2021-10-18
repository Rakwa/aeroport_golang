package subscribers

import (
	"broker/broker"
	"broker/config"
	"context"
	"flag"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var db *mongo.Client

func connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func RunSub(subscriberType string, subscriberName string) {
	topic := "airport"
	flag.Parse()
	dbClient, _, _, _ := connect(config.AppConfig.DbUrl)
	db = dbClient
	var client = broker.Connect(subscriberName, "")
	if subscriberType == "db" {
		fmt.Println("db init")
		client.Subscribe(topic, 1, SeedDb)
	} else {
		client.Subscribe(topic, 1, CsvFileExport)
	}
	for {
	}
}
