package main

import (
	"broker/config"
	"broker/publishers"
	"broker/subscribers"
)

func main() {
	// runType := flag.String("exec", "sub", "exec type")
	// if *runType == "sub" {
	// 	subscribers.RunSub()
	// }
	config.ReadConfig()
	go publishers.RunPub()
	go subscribers.RunSub()
	//
	for {
	}
}
