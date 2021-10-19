package main

import (
	"api/helper"
	"api/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

var collection = helper.ConnectDB()

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/airports/{id}/measures", getMeasures).Methods("GET")
	router.HandleFunc("/api/airports/{id}/averages", getAverages).Methods("GET")

	fmt.Println("API availabe in localhost:3333")

	log.Fatal(http.ListenAndServe(":3333", router))

}

func getMeasures(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var measures []models.Measure

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

	if startDate != "" {
		startDateConvert, err := time.Parse(time.RFC3339, startDate)
		if err != nil {
			helper.GetError(err, writer)
			return
		}
		startDateUnix = startDateConvert.Unix()
	}

	if endDate != "" {
		endDateConvert, err := time.Parse(time.RFC3339, endDate)
		if err != nil {
			helper.GetError(err, writer)
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

	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, writer)
		return
	}

	defer cur.Close(context.TODO())

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

func getAverages(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(request)
	var date = request.URL.Query().Get("date")

	filter := bson.M{
		"airportid": params["id"],
	}

	dateStartUnix := int64(-1)
	dateEndUnix := int64(-1)

	if date != "" {
		dateConvert, err := time.Parse("2006-01-02", date)
		fmt.Println(dateConvert)
		if err != nil {
			helper.GetError(err, writer)
			return
		}
		dateStartUnix = dateConvert.Unix()
		fmt.Println(dateConvert)
		dateEndUnix = dateConvert.Add(24*time.Hour - 1).Unix()
	}

	if dateStartUnix != -1 && dateEndUnix != -1 {
		filter["date"] = bson.M{"$gte": dateStartUnix, "$lte": dateEndUnix}
	}

	matchStage := bson.D{{"$match", filter}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$type"}, {"average", bson.D{{"$avg", "$value"}}}}}}

	cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{matchStage, groupStage})

	if err != nil {
		helper.GetError(err, writer)
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
