# Projet Aéroport - Broker

Broker permettant de faire la liaison entre les capteurs et la base de données

## Installation

Nécessite MongoDB et Mosquitto

Pour lancer MongoDB et Mosquitto via docker :
```shell
docker-compose up -d
```

- MongoDB : [localhost:27017](localhost:27017)
- Mosquitto : [localhost:52883](localhost:52883)

Lancer un capteur (type possible : temperature / wind / pressure) : 
```shell
go run broker -type=pressure -sensorId=capteur3 -airportId=NTA
```