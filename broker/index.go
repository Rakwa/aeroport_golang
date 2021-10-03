package main

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gomodule/redigo/redis"
)

var conn, err = redis.Dial("tcp", "localhost:6379")

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Println(conn)
	ok, err := conn.Do("RPUSH", "TEMP", msg.Payload())
	fmt.Println(ok)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Message %s received on topic %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectionLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection Lost: %s\n", err.Error())
}

func main() {

	options := mqtt.NewClientOptions()
	options.AddBroker("tcp://127.0.0.1:1884")
	options.SetClientID("<frfre")
	client := mqtt.NewClient(options)
	token := client.Connect()
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectionLostHandler

	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	topic := "a/b/a"
	token = client.Subscribe(topic, 1, messagePubHandler)
	token.Wait()
	fmt.Printf("Subscribed to topic %s\n", topic)

	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("%d", i)
		token = client.Publish(topic, 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}

	client.Disconnect(100)
}
