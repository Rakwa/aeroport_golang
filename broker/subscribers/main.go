package subscribers

import (
	"broker/broker"
	"broker/config"
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

var db *mongo.Client

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

func SeedDb(client mqtt.Client, msg mqtt.Message) {
	fmt.Println("MESSAGE RECU")
	fmt.Println(string(msg.Payload()))
	data := publishers.Data{}
	json.Unmarshal([]byte(msg.Payload()), &data)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := db.Database("go_project").Collection("aeroports")

	/*
	   Insert documents
	*/

	_, insertErr := collection.InsertOne(ctx, data)
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	/*
	   Iterate a cursor and print it
	*/

}
func sub(dbClient *mongo.Client, ctx context.Context, mqttClient mqtt.Client) {
	token := mqttClient.Subscribe("airport", 1, SeedDb)
	token.Wait()
	fmt.Printf("Subscribed to topic %s", "airport")
}

func RunSub(subscriberType string, subscriberName string) {
	topic := "airport"
	flag.Parse()
	var client = broker.Connect(subscriberName, "")
	if subscriberType == "db" {
		client.Subscribe(topic, 1, SeedDb)
	} else {
		client.Subscribe(topic, 1, CsvFileExport)
	}
	fmt.Println("db init")
	//sub(client)
  for { }
}
