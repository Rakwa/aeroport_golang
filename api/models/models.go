package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Measure struct {
	Id        primitive.ObjectID `json:"_id"`
	SensorId  string             `json:"sensorid"`
	AirportId string             `json:"airportid"`
	Date      int                `json:"date"`
	Type      string             `json:"type"`
	Value     float64            `json:"value"`
}
