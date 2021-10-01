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

## Lancement

Pour lancer MongoDB via docker : 
```shell
docker-compose up -d
```