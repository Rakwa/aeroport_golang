# Projet Aéroport - Publishers / Subscribers ✈️

Gestion des capteurs et des enregistrement en base de données et fichiers csv

## Installation avec Docker (recommandée) 🐳

[Voir la documentation générale](../README.md)

## Installation sans Docker :warning:

1. Installer MongoDB sur votre machine
2. Installer Mosquitto sur votre machine
3. Modifier la configuration des publishers et subscribers si les ports de MongoDB et Mosquitto sont différents (optionnel)
4. Compiler le code des publishers et subscribers (optionnel)
```shell
go build
```
5. Lancer les publishers et subscribers

```shell
# Capteur de temperature
./brokerApp -type=temperature -sensorId=capteur1 -airportId=NTEs

# Capteur de pression
./brokerApp -type=pressure -sensorId=capteur2 -airportId=NTEs

# Capteur de vent
./brokerApp -type=wind -sensorId=capteur3 -airportId=NTEs

# Enregistreur en base de données
./brokerApp -mode=subscriber -subName=db -subType=db

# Enregistreur dans un fichier
./brokerApp -mode=subscriber -subName=csv -subType=csv
```
:warning: Utiliser un id (sensorId ou subName) différent à chaque fois 

## Script de test
Un script de test composé de 10 capteurs avec 3 aéroports, un enregistrement en base de données et dans un fichier est disponible

```shell
# Pour lancer le script de test
./docker.sh
```