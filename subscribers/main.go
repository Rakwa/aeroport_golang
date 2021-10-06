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
func Connect(airportId string) mqtt.Client {
	var broker = "broker.emqx.io"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	opts.OnConnect = connectHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	sub(client)
	return client
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
func sub(client mqtt.Client) {
	topic := "topic/test"
	token := client.Subscribe(topic, 1, SeedDb)
	token.Wait()
	fmt.Printf("Subscribed to topic %s", topic)
}
func main() {
	StartDb()
	fmt.Println("db init")
	Connect("NTA")

	for {

	}
}
