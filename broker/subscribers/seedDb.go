package subscribers

import (
	"broker/config"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/bson"
)

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
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := config.Client.Database("blog").Collection("posts")

	/*
	   Insert documents
	*/
	docs := []interface{}{
		bson.D{{"title", "World"}, {"body", "Hello World"}},
		bson.D{{"title", "Mars"}, {"body", "Hello Mars"}},
		bson.D{{"title", "Pluto"}, {"body", "Hello Pluto"}},
	}

	res, insertErr := collection.InsertMany(ctx, docs)
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	fmt.Println(res)
	/*
	   Iterate a cursor and print it
	*/
	cur, currErr := collection.Find(ctx, bson.D{})

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)

	var posts []Post
	if err = cur.All(ctx, &posts); err != nil {
		panic(err)
	}
	fmt.Println(posts)
}
