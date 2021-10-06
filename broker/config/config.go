package config

/*
	General configuration with config.yaml
*/
import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v2"
)

const configPath = "config/config.yaml"

type Config struct {
	MqttURL             string        `yaml:"mqttURL"`
	PublishersFrequency time.Duration `yaml:"publishersFrequency"`
	//add parameter here when new parameter add in config.yaml
}

var AppConfig *Config
var Client *mongo.Client

/**
Transform config.yaml to Config structure
*/
func ReadConfig() {
	f, err := os.Open(configPath)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&AppConfig)

	if err != nil {
		fmt.Println(err)
	}
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
