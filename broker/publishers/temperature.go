package publishers

import (
	"math"
	"math/rand"
	"time"
)

/*
  Calculation of the next temperature sensor value
  value : previous value
  return next value
*/
func nextValueTemperatureSensor(value float64) float64 {
	return math.Round((value+0.1*float64(rand.Intn(3)-1))*100) / 100
}

/*
	Create new temperature sensor
	sensorId
	airportId
*/
func CreateTemperatureSensor(sensorId string, airportId string) {
	rand.Seed(time.Now().UnixNano())
	value := Data{sensorId, airportId, time.Now().Unix(), "temperature", float64(rand.Intn(30))}
	CreateSensor(value, nextValueTemperatureSensor)
}
