package subscribers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"strconv"
	"time"
)

func CsvFileExport(client mqtt.Client, msg mqtt.Message) {
	data := Data{}
	json.Unmarshal([]byte(msg.Payload()), &data)
	if _, err := os.Stat("../csv/"); os.IsNotExist(err) {
		err = os.Mkdir("../csv/", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	measure_date := time.Unix(data.Date, 0)

	name := data.AirportId + "_" + strconv.Itoa(measure_date.Year()) + "-" + strconv.Itoa(int(measure_date.Month())) + "-" + strconv.Itoa(measure_date.Day()) + ".csv"
	record := [][]string{{"sensor_id", "airport_id", "date", "type", "value"}, {data.SensorId, data.AirportId, measure_date.String(), data.Type, fmt.Sprintf("%f", data.Value)}}
	if _, err := os.Stat("../csv/" + name); os.IsNotExist(err) {
		os.Create("../csv/" + name)
	} else {
		record = record[1:]
	}
	f, _ := os.OpenFile("../csv/"+name, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	w := csv.NewWriter(f)
	if err2 := w.WriteAll(record); err2 != nil {
		log.Fatalln("error writing record to file", err2)
	}
}
