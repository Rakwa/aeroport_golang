package controller

import (
	"api/helper"
	"api/models"
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

// GetMeasures godoc
// @Summary Get measures (temperature, pressure, wind) in a airport between two date
// @Description Get measures (temperature, pressure, wind) in a airport between two date
// @Tags Airports
// @Accept  json
// @Produce  json
// @Param id path string true "id of the airport (example : NTE)"
// @Param type query string false "sensor type (temperature, wind, pressure)"
// @Param startDate query string false "start date (example : 2021-04-04T22:08:41Z)"
// @Param endDate query string false "end date (example : 2021-04-04T22:08:41Z)"
// @Success 200 {array} models.Measure
// @Failure 400 {object} helper.ErrorResponse
// @Router /airports/{id}/measures [get]
func GetMeasures(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var measures []models.Measure

	//Format parameters
	var params = mux.Vars(request)
	var sensorType = request.URL.Query().Get("type")
	var startDate = request.URL.Query().Get("startDate")
	var endDate = request.URL.Query().Get("endDate")

	filter := bson.M{
		"airportid": params["id"],
	}

	if sensorType != "" {
		filter["type"] = sensorType
	}

	startDateUnix := int64(-1)
	endDateUnix := int64(-1)

	errorDateFormat := errors.New("Invalid date format (correct format RFC3339 : 2021-04-04T22:08:41Z)")

	if startDate != "" {
		startDateConvert, err := time.Parse(time.RFC3339, startDate)
		if err != nil {
			helper.GetError(errorDateFormat, writer, 400)
			return
		}
		startDateUnix = startDateConvert.Unix()
	}

	if endDate != "" {
		endDateConvert, err := time.Parse(time.RFC3339, endDate)
		if err != nil {
			helper.GetError(errorDateFormat, writer, 400)
			return
		}
		endDateUnix = endDateConvert.Unix()
	}

	if startDateUnix != -1 && endDateUnix != -1 {
		filter["date"] = bson.M{"$gte": startDateUnix, "$lte": endDateUnix}
	} else if startDateUnix != -1 {
		filter["date"] = bson.M{"$gte": startDateUnix}
	} else if endDateUnix != -1 {
		filter["date"] = bson.M{"$lte": endDateUnix}
	}

	//Get data
	cur, err := helper.Collection.Find(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, writer, 500)
		return
	}

	defer cur.Close(context.TODO())

	//Format response
	for cur.Next(context.TODO()) {
		var measure models.Measure
		err := cur.Decode(&measure)
		if err != nil {
			log.Fatal(err)
		}

		measures = append(measures, measure)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(writer).Encode(measures)
}

// GetAverages godoc
// @Summary Get averages of measures (temperature, pressure, wind) in a airport
// @Description Get averages of measures (temperature, pressure, wind) in a airport
// @Tags Airports
// @Accept  json
// @Produce  json
// @Param id path string true "id of the airport (example : NTE)"
// @Param date query string false "date (example : 2021-04-04)"
// @Success 200 {array} models.Average
// @Failure 400 {object} helper.ErrorResponse
// @Router /airports/{id}/averages [get]
func GetAverages(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	//Format parameters
	var params = mux.Vars(request)
	var date = request.URL.Query().Get("date")

	filter := bson.M{
		"airportid": params["id"],
	}

	dateStartUnix := int64(-1)
	dateEndUnix := int64(-1)

	if date != "" {
		dateConvert, err := time.Parse("2006-01-02", date)

		errorDateFormat := errors.New("Invalid date format (correct format : 2021-04-04)")

		if err != nil {
			helper.GetError(errorDateFormat, writer, 400)
			return
		}
		dateStartUnix = dateConvert.Unix()
		dateEndUnix = dateConvert.Add(24*time.Hour - 1).Unix()
	}

	if dateStartUnix != -1 && dateEndUnix != -1 {
		filter["date"] = bson.M{"$gte": dateStartUnix, "$lte": dateEndUnix}
	}

	//Get data
	matchStage := bson.D{{"$match", filter}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$type"}, {"average", bson.D{{"$avg", "$value"}}}}}}

	cur, err := helper.Collection.Aggregate(context.TODO(), mongo.Pipeline{matchStage, groupStage})

	if err != nil {
		helper.GetError(err, writer, 500)
		return
	}

	defer cur.Close(context.TODO())

	var showsWithInfo []bson.M

	err = cur.All(context.TODO(), &showsWithInfo)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(writer).Encode(showsWithInfo)

}

// GetAirports godoc
// @Summary Get list of airports in the database
// @Description Get list of airports in the database
// @Tags Airports
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Airport
// @Failure 400 {object} helper.ErrorResponse
// @Router /airports [get]
func GetAirports(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	//Get data
	matchStage := bson.D{{"$match", bson.D{}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$airportid"}}}}

	cur, err := helper.Collection.Aggregate(context.TODO(), mongo.Pipeline{matchStage, groupStage})

	if err != nil {
		helper.GetError(err, writer, 500)
		return
	}

	defer cur.Close(context.TODO())

	var airports []models.Airport

	//Format response
	for cur.Next(context.TODO()) {
		var airport models.Airport
		err := cur.Decode(&airport)
		if err != nil {
			log.Fatal(err)
		}

		airport = helper.AirportsData[airport.Id]
		if airport.Id != "" {
			airports = append(airports, helper.AirportsData[airport.Id])
		}
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(writer).Encode(airports)

}
