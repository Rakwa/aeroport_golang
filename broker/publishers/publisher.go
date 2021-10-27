package publishers

import (
	"brokerApp/broker"
	"brokerApp/config"
	"encoding/json"
	"fmt"
	"time"
)

/*
  Structure of data send to the broker
*/
type Data struct {
	SensorId  string  `json:"sensor_id"`
	AirportId string  `json:"airport_id"`
	Date      int64   `json:"date"`
	Type      string  `json:"type"`
	Value     float64 `json:"value"`
}

/*
	Create sensor connected to the broker
	initData : default value
	nextValue : pointer of function to determine next value
*/
func CreateSensor(initData Data, nextValue func(value float64) float64) {
	client := broker.Connect(initData.SensorId, initData.AirportId)

	value := initData
	for {
		value.Date = time.Now().Unix()
		value.Value = nextValue(value.Value)
		jsonValue, err := json.Marshal(value)

		if err != nil {
			fmt.Println("Error:", err)
		}

		client.Publish("airport", config.AppConfig.QOS, false, string(jsonValue))
		time.Sleep(time.Second * config.AppConfig.PublishersFrequency)
	}
}

/**
  Create new sensor
  command : go run broker -type=pressure -sensorId=capteur3 -airportId=NTA
  Types available : pressure, wind, temperature
*/
func RunPub(sensorType string, sensorId string, airportId string) {
	sensorTypeValue := sensorType
	sensorIdValue := sensorId
	airportIdValue := airportId

	if sensorTypeValue == "temperature" {
		fmt.Println("Temperature sensor (" + sensorIdValue + ") initialisation for airport " + airportIdValue)
		CreateTemperatureSensor(sensorIdValue, airportIdValue)
	} else if sensorTypeValue == "wind" {
		fmt.Println("Wind sensor (" + sensorIdValue + ") initialisation for airport " + airportIdValue)
		CreateWindSensor(sensorIdValue, airportIdValue)
	} else if sensorTypeValue == "pressure" {
		fmt.Println("Pressure sensor (" + sensorIdValue + ") initialisation for airport " + airportIdValue)
		CreatePressureSensor(sensorIdValue, airportIdValue)
	} else {
		fmt.Println("Unknown sensor type")
	}
}
