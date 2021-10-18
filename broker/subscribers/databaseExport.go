package subscribers

import (
	"broker/publishers"
	"context"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

/*
	Export data in Mongo db
*/
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
