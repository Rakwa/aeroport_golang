package broker

/*
	File to configure broker
*/
import (
	"broker/config"
	"broker/subscribers"
	"fmt"
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
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", config.AppConfig.MqttURL))
	opts.SetClientID(sensorId)
	opts.OnConnect = connectHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	client.Subscribe(airportId, 1, subscribers.CsvFileExport)
	return client
}
