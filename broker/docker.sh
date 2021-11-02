#!/bin/bash

#subscribers
./brokerApp -mode=subscriber -subName=db -subType=db &
./brokerApp -mode=subscriber -subName=csv -subType=csv &

#publishers

#NTE
./brokerApp -type=pressure -sensorId=pressureNTE -airportId=NTE &
./brokerApp -type=wind -sensorId=windNTE -airportId=NTE &
./brokerApp -type=temperature -sensorId=temperatureNTE1 -airportId=NTE &
./brokerApp -type=temperature -sensorId=temperatureNTE2 -airportId=NTE &

# ORY
./brokerApp -type=pressure -sensorId=pressureORY -airportId=ORY &
./brokerApp -type=wind -sensorId=windORY -airportId=ORY &
./brokerApp -type=temperature -sensorId=temperatureORY -airportId=ORY &

# LYS
./brokerApp -type=pressure -sensorId=pressureLYS -airportId=LYS &
./brokerApp -type=wind -sensorId=windLYS -airportId=LYS &
./brokerApp -type=temperature -sensorId=temperatureLYS -airportId=LYS


