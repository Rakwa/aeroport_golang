package broker

/*
	File to configure broker
*/
import (
	"broker/config"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

/*
	On connection
*/
func connectHandler(client mqtt.Client) {
	fmt.Println("coucou")
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
	fmt.Printf(sensorId + airportId)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", config.AppConfig.MqttURL))
	opts.SetClientID(sensorId + airportId)
	opts.OnConnect = connectHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return client
}
