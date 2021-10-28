# Projet Aéroport - API ✈️

API pour récupérer des informations sur les aéroports avec les mesures des différents capteurs

## Installation avec Docker (recommandée) 🐳

[Voir la documentation générale](../README.md)

## Installation sans Docker :warning:

1. Installer MongoDB sur votre machine
2. Lancer le script des publishers et subscribers : [voir documentation](../broker/README.md)
3. Modifier la configuration de l'API si le port de MongoDB est différent (optionnel)
7. Lancer l'API
```shell
go run main.go
```

- L'API est accessible à l'adresse : [localhost:3333](localhost:3333)
- La documentation se situe à l'adresse : [localhost:3333](localhost:3333)

## Autres commandes
Pour générer et mettre à jour la documentation Swagger :
```shell
swag init
```
