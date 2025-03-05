# POKEDEX CLI

![Demo](/demo_pokedex-cli.gif)

Explorez le monde des Pokemons et collectionnez-les, directement depuis votre terminal !

Cette application en ligne de commande utilise l'API [Pokeapi](https://pokeapi.co/api/v2) pour récupérer des informations sur les Pokemon et les lieux. 

## Table des matières

- [Installation](#installation)
- [Liste des commandes](#liste-des-commandes)
- [Contexte autour développement](#features)

## Installation

> Nécessaire d'installer Go pour complier l'exécutable 

```bash
# Clone le dépôt distant
git clone https://github.com/benKapl/pokedex-cli.git

# Déplacez vous dans le dossier
cd pokedex-cli

# Lancer la compilation
go build
```

## Liste des commandes


Voici la liste des commandes disponibles dans le Pokédex :
> A noter: l'interface est écrite en anglais

- **exit** : Quitter le Pokédex.
- **help** : Afficher la liste et les détails des commandes disponibles.
- **map** : Obtenir la page suivante des localisations.
- **mapb** : Obtenir la page précédente des localisations.
- **explore <nom_du_lieu>** : Lister les Pokémon présents dans la zone spécifiée.
- **catch <nom_du_pokemon>** : Lancer une Pokéball pour capturer un Pokémon !
- **inspect <nom_du_pokemon>** : Afficher des informations sur un Pokémon capturé.
- **pokedex** : Afficher la liste de tous les Pokémon capturés.

## Contexte autour du développement

Le CLI Pokedex est un projet guidé de la plateforme [Boot.dev](https://boot.dev) pour l'apprentissage de Go et des clients HTTP.

L'application permet de mettre en pratiques les notions suivantes : 
- Parsing de JSON en Go
- Réalisation de requêtes HTTP en Go
- Création d'un outil CLI pour faciliter l'interaction avec un serveur back-end
- Développement local en Go et l'utilisation des outils associés
- Mise en cache pour améliorer les performances
- Utilisation de la librarie de tests Go

