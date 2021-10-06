package main

import (
	"broker/config"
	"broker/publishers"
	"flag"
	"fmt"
)

/**
Create new sensor

command : go run broker -type=pressure -sensorId=capteur3 -airportId=NTA
Types available : pressure, wind, temperature
*/
func main() {
	sensorType := flag.String("type", "temperature", "Sensor type")
	sensorId := flag.String("sensorId", "sensorDefault", "Sensor id")
	airportId := flag.String("airportId", "NTE", "Airport id")
	flag.Parse()
	config.StartDb()
	config.ReadConfig()
	sensorTypeValue := *sensorType
	sensorIdValue := *sensorId
	airportIdValue := *airportId

	if sensorTypeValue == "temperature" {
		fmt.Println("Temperature sensor (" + sensorIdValue + ") initialisation for airport " + airportIdValue)
		publishers.CreateTemperatureSensor(sensorIdValue, airportIdValue)
	} else if sensorTypeValue == "wind" {
		fmt.Println("Wind sensor (" + sensorIdValue + ") initialisation for airport " + airportIdValue)
		publishers.CreateWindSensor(sensorIdValue, airportIdValue)
	} else if sensorTypeValue == "pressure" {
		fmt.Println("Pressure sensor (" + sensorIdValue + ") initialisation for airport " + airportIdValue)
		publishers.CreatePressureSensor(sensorIdValue, airportIdValue)
	} else {
		fmt.Println("Unknown sensor type")
	}
}
