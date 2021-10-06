package publishers

import (
	"broker/broker"
	"broker/config"
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

		client.Publish(initData.AirportId, 2, false, string(jsonValue))
		time.Sleep(time.Second * config.AppConfig.PublishersFrequency)
	}
}
