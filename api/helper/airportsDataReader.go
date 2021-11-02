package helper

import (
	"api/models"
	"encoding/csv"
	"log"
	"os"
)

func createAirportsList(data [][]string) map[string]models.Airport {
	var airportsList = make(map[string]models.Airport)
	for i, line := range data {
		if i > 0 { // omit header line
			var rec models.Airport
			for j, field := range line {
				if j == 2 {
					rec.Name = field
				} else if j == 7 {
					rec.City = field
				} else if j == 9 {
					rec.Id = field
				}
			}
			if rec.Id != "" {
				airportsList[rec.Id] = rec
			}
		}
	}
	return airportsList
}

func GetAirportsData() map[string]models.Airport {
	f, err := os.Open("airports.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// convert records to array of structs
	airportsList := createAirportsList(data)

	return airportsList
}

var AirportsData = GetAirportsData()
