# Installation de l'environnement

## Installation de Go

- Téléchargez Go 1.16 pour votre environnement : [https://golang.org/dl/](https://golang.org/dl/)
- Suivez les instructions d’installation : [https://golang.org/doc/install](https://golang.org/doc/install)
- Vérifiez l’installation dans un terminal :

```shell
$ go version
```

- Ajoutez la variable d’environnement GOPATH pour pointer sur votre workspace. Par exemple :

```shell
$ GOPATH=$HOME\go
$ GOPATH=%USERPROFILE%\go
```

<!-- .element: class="big-code" -->

Notes:
Normalement, l’installateur ajoute Go au PATH, mais si ce n’est pas le cas, il faut le faire manuellement.
GOPATH doit contenir le chemin vers le répertoire de travail Go (là où l’on va trouver les répertoires src, bin, etc.)

FSA :
Install normalement déjà faite par les participants, mais pour ceux qui ne l'auront pas fait c'est le moment.
On pourra éventuellement parler de https://github.com/moovweb/gvm le version manager pour go. Permet d'installer plusieurs version sur le poste de travail.

##==##

# Installation de l'environnement

## Les commandes de Go

- **build** compile packages et dépendances
- **fmt** lance gofmt sur les sources
- **test** lance les tests
- **tool** lance les outils spécifiques (tracer, etc...)
- **get** télécharge et installe les packages and dépendances
- **mod** gestion des dépendances

##==##

# Installation de l'environnement

## Récupération des sources

- Votre workspace de base contient

  - **bin :** binaires de vos appli
  - **pkg :** vos objets à linker
  - **src :** toutes vos sources
  - Ajouter les bin Go à votre path : `PATH=$PATH:$GOPATH/bin`

- Il vous faut cloner :

  - git clone <a href="https://github.com/sfeir-open-source/sfeir-school-go">https://github.com/sfeir-open-source/sfeir-school-go</a>
  - **cd** dans **sfeir-school-go/go-100**

##==##

# Installation de l'environnement

## Installation de Visual Studio Code

- Télé-chargez et installez VSC pour votre environnement: [https://code.visualstudio.com](https://code.visualstudio.com/)
- Installez le plugin Go: <img src="./assets/go-100/images/plugins.jpg">, puis chercher “go”.

- En ouvrant un fichier **.go**, le message _“Analysis Tools Missing"_ apparaîtra : cliquez dessus pour compléter l’installation automatiquement.
- Il est conseillé d’activer la sauvegarde automatique.
