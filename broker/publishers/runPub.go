package publishers

import (
	"fmt"
)

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
