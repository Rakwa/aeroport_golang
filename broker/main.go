package main

import (
	"broker/config"
	"broker/publishers"
	"broker/subscribers"
	"flag"
)

func main() {
	config.ReadConfig()
	mode := flag.String("mode", "publisher", "Subscriber or publisher ?")
	subscriberType := flag.String("subType", "db", "db or csv ?")
	subscriberName := flag.String("subName", "my_sub_name", "subscriber name")
	sensorType := flag.String("type", "temperature", "Sensor type")
	sensorId := flag.String("sensorId", "sensorDefault", "Sensor id")
	airportId := flag.String("airportId", "NTE", "Airport id")
	flag.Parse()
	if *mode == "publisher" {
		publishers.RunPub(*sensorType, *sensorId, *airportId)
	} else {
		subscribers.RunSub(*subscriberType, *subscriberName)
	}
	for {
	}
}
