package helper

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

/*
	Connect to MongoDB database
*/
func ConnectDB() *mongo.Collection {

	ReadConfig()
	clientOptions := options.Client().ApplyURI(AppConfig.DbUrl)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("go_project").Collection("aeroports")
	return collection
}

var Collection = ConnectDB()
