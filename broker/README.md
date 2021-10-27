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

Lancer l'application avec les capteurs de test :
```shell
./docker.sh
```
Lancer un capteur (type possible : temperature / wind / pressure) : 
```shell
./brokerApp -type=pressure -sensorId=capteur4 -airportId=NTEs
```
Lancer un subscriber (type possible : csv / db) :
```shell
./brokerApp -mode=subscriber -subName=colorgitano -subType=csv
```

### Autres commandes
Compiler l'application
```shell
go build
```