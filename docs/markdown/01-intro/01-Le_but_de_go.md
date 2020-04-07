
# Avant de se lancer

## Le but de Go

- Rassembler les bonnes idées provenant d’autres langages,
- Abolir les fonctionnalités débouchant sur un code complexe et peu fiable…

...dans le but d’écrire du code simple, expressif, robuste et efficace.

Notes:
- Principalement inspiré du C.
- Particulièrement bénéfique sur de gros projets, avec de grandes équipes.

##==##

# Avant de se lancer

## Un langage généraliste
- Comme le C, il sera adapté dans pratiquement tous les domaines de programmation.
- Idéal pour le cloud.
- Déjà utilisé dans le graphisme, les applications mobiles, le machine learning, WASM, ...


##==##

# Avant de se lancer

Déjà adopté par des grands
- Google
- Docker
- Kubernetes
- Dropbox
- Spotify
- Hashicorp
- SoundCloud
- etc...


##==##

# Le language

## Déjà adopté par des grands

- Né en 2009 chez Google (après les processeurs multi-coeurs) et OSS
- Binaire compilé autoporteur (début plugin depuis Go 1.8)
- Orienté objet
- Garbage collector (sub millisecond pour 17 Go de heap)
- Pointeurs 😱
- Goroutines
  - Assimilable à un thread
  - Mais ce n’est **PAS** un thread ⇒ **beaucoup plus léger**

- Channels
- **Do not communicate by sharing memory; share memory by communicating.**
Synchronisation
Multiplexage (**select**)


Notes:
- OFU
- CSP (1977) Communicating sequential processes

##==##

# Le language

## Les mots clés
- **Dépendances :** import package
- **Conditionnelles :** if else switch case fallthrough break default goto select
- **Itérations :** for range continue
- **Type :** var func interface struct chan const type map make
- **Misc :** defer go return panic recover

![h-350](./assets/images/mots_clés.JPG)<!-- .element: class="special-Intro-01-le-but-de-go-bottom-image" -->
![h-350](./assets/images/i_know.JPG)<!-- .element: class="special-Intro-01-le-but-de-go-bottom-image" -->
Notes:
- OFU
- moins de 30 mots-clés
- Public/private (exporté/non exporté) => Majuscule/minuscule







