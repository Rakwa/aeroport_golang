package publishers

import (
	"math/rand"
	"time"
)

/*
  Calculation of the next wind sensor value
  value : previous value
  return next value
*/
func nextValueWindSensor(value float64) float64 {
	return value + float64(rand.Intn(3)-1)
}

/*
	Create new wind sensor
	sensorId
	airportId
*/
func CreateWindSensor(sensorId string, airportId string) {
	rand.Seed(time.Now().UnixNano())
	value := Data{sensorId, airportId, time.Now().Unix(), "wind", float64(rand.Intn(80) + 20)}
	CreateSensor(value, nextValueWindSensor)
}
