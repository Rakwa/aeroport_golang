package subscribers

import (
	"broker/broker"
	"broker/publishers"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
	On connection
*/
func connectHandler(client mqtt.Client) {
	fmt.Println("Connected")
}

/*
	On connection lost
*/
func connectionLostHandler(client mqtt.Client, err error) {
	fmt.Printf("Connection Lost: %s\n", err.Error())
}

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

func SeedDb(client mqtt.Client, msg mqtt.Message) {
	fmt.Println(string(msg.Payload()))
	data := publishers.Data{}
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
func sub(client mqtt.Client) {
	topic := "topic/test"
	token := client.Subscribe(topic, 1, SeedDb)
	token.Wait()
	fmt.Printf("Subscribed to topic %s", topic)
}

func RunSub(subscriber_type string, subscriber_name string) {
	topic := "airport"
	flag.Parse()
	var client = broker.Connect(subscriber_name, "")
	if subscriber_type == "db" {
		client.Subscribe(topic, 1, SeedDb)
	} else {
		client.Subscribe(topic, 1, CsvFileExport)
	}
	fmt.Println("db init")
	//sub(client)
	for {

	}
}
