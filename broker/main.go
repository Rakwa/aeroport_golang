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
	subscriber_type := flag.String("subType", "db", "db or csv ?")
	subscriber_name := flag.String("subName", "my_sub_name", "subscriber name")
	sensorType := flag.String("type", "temperature", "Sensor type")
	sensorId := flag.String("sensorId", "sensorDefault", "Sensor id")
	airportId := flag.String("airportId", "NTE", "Airport id")
	flag.Parse()
	if *mode == "publisher" {
		publishers.RunPub(*sensorType, *sensorId, *airportId)
	} else {
		subscribers.RunSub(*subscriber_type, *subscriber_name)
	}
	for {
	}
}
