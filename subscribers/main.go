package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func StartDb() {
	Client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer Client.Disconnect(ctx)
}

type Data struct {
	SensorId  string  `json:"sensor_id"`
	AirportId string  `json:"airport_id"`
	Date      string  `json:"date"`
	Type      string  `json:"type"`
	Value     float64 `json:"value"`
}

func SeedDb(client mqtt.Client, msg mqtt.Message) {
	fmt.Println(string(msg.Payload()))
	data := Data{}
	json.Unmarshal([]byte(msg.Payload()), &data)
	fmt.Printf(string(data.SensorId))
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// collection := config.Client.Database("go_project").Collection("aeroports")

	// /*
	//    Insert documents
	// */

	// res, insertErr := collection.InsertMany(ctx, docs)
	// if insertErr != nil {
	// 	log.Fatal(insertErr)
	// }
	// fmt.Println(res)
	// /*
	//    Iterate a cursor and print it
	// */
	// cur, currErr := collection.Find(ctx, bson.D{})

	// if currErr != nil {
	// 	panic(currErr)
	// }
	// defer cur.Close(ctx)

	// var posts []Data
	// fmt.Println(posts)
}
func main() {

}
