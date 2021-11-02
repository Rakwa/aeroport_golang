package publishers

import (
	"math"
	"math/rand"
	"time"
)

/*
  Calculation of the next pressure sensor value
  value : previous value
  return next value
*/
func nextValuePressureSensor(value float64) float64 {
	return math.Round((value+0.1*float64(rand.Intn(3)-1))*100) / 100
}

/*
	Create new pressure sensor
	sensorId
	airportId
*/
func CreatePressureSensor(sensorId string, airportId string) {
	rand.Seed(time.Now().UnixNano())
	value := Data{sensorId, airportId, time.Now().Unix(), "pressure", float64(rand.Intn(30) + 1000)}
	CreateSensor(value, nextValuePressureSensor)
}
