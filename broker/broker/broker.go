package broker

/*
	File to configure broker
*/
import (
	"broker/config"
	"broker/subscribers"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

/*
	On message published
*/
func messagePubHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Message %s received on topic %s\n", msg.Payload(), msg.Topic())
}

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

/*
	Connect publisher
	clientId
	airportId //remove after add mongoDB
*/
func Connect(sensorId string, airportId string) mqtt.Client {
	options := mqtt.NewClientOptions()
	options.AddBroker("tcp://" + config.AppConfig.MqttURL)
	options.SetClientID(sensorId)
	client := mqtt.NewClient(options)
	token := client.Connect()
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectionLostHandler

	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	token = client.Subscribe(airportId, 1, subscribers.SeedDb)

	return client
}
