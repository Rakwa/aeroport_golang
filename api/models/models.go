package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Measure struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	SensorId  string             `json:"sensorid"`
	AirportId string             `json:"airportid"`
	Date      int                `json:"date"`
	Type      string             `json:"type"`
	Value     float64            `json:"value"`
}

type Average struct {
	Id      string `json:"_id"`
	Average string `json:"average"`
}

type Airport struct {
	Id   string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name"`
	City string `json:"city"`
}
