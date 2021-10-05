# Projet Aéroport

Système de collecte et de restitution de données météo des aéroports (température, vitesse du vent, pression atmosphérique)

Projet réalisé dans le cadre du module "Infrastructures d'intégration" de l'IMT Atlantique 

Créé par Théo LETOUZÉ, Clément NICOLAS, Loïs GIGAUD, Julien RAQUOIS

## Contenu

- 3 publishers en Go envoyant les données des capteurs
- Un broker MQTT Mosquitto réalisant la liaison entre les publishers et la base de données et le fichier csv
- Une base de données MongoDB
- Une API REST en ...
- Un client REST en VueJS (v3)

## Installation

MongoDB et Mosquitto fonctionne via une image docker

Pour lancer MongoDB et Mosquitto via docker : 
```shell
docker-compose up -d
```

- MongoDB : [localhost:27017](localhost:27017)
- Mosquitto : [localhost:52883](localhost:52883)

## Broker

[Documentation complète sur le broker](broker/README.md)

Pour lancer un capteur depuis le dossier broker (type possible : temperature / wind / pressure) :
```shell
go run broker -type=pressure -sensorId=capteur3 -airportId=NTA
```