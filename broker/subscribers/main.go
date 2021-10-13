package subscribers

import (
	"broker/broker"
	"broker/config"
	"context"
	"encoding/json"
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

type Data struct {
	SensorId  string  `json:"sensor_id"`
	AirportId string  `json:"airport_id"`
	Date      string  `json:"date"`
	Type      string  `json:"type"`
	Value     float64 `json:"value"`
}

func SeedDb(client mqtt.Client, msg mqtt.Message) {
	fmt.Println("MESSAGE RECU")
	fmt.Println(string(msg.Payload()))
	data := Data{}
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
func RunSub() {
	dbClient, ctx, _, _ := connect(config.AppConfig.DbUrl)
	db = dbClient
	var client = broker.Connect("coucou")
	sub(dbClient, ctx, client)

}
