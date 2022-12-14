# Orienté Objet - 10

## Les Generics

Le but des types et fonctions génériques est de permettre de définir :

 * Des caractéristiques spéciales de structures de données (liste chaînée, comparable, etc.)
 * Des actions réutilisables sur des types particuliers (tri, filtre, etc.)

Par crainte de dégrader les performances ou complexifier le langage
 [les génériques](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#operations-based-on-type-sets) étaient dédiés initialement à Go 2.0
mais sont apparus avec la version 1.18 (15/03/2022)

##==##

# Orienté Objet - 10

## Les fonctions génériques

<!-- .slide: class="with-code" -->

Une première version des génériques est déjà possible grâces aux `interface{}`.

Mais les types et fonctions génériques vont un cran plus loin.

```Go
func Index[T comparable](s []T, x T) int
```

Les fonctions génériques permettent de fonctionner avec plusieurs type de paramètres.

Les types apparaissent `[]` avant la déclaration des paramètres de la fonction.

Le type `comparable` est une contrainte qui permet d'utiliser les opérateur `==` et `!=`
sur les valeurs des types.

##==##

# Orienté Objet - 10

## Les fonctions génériques

<!-- .slide: class="with-code" -->

Il existe d'autres contraintes et on peut créer ses propres contraintes.

```Go
// Ordered is a type constraint that matches any ordered type.
// An ordered type is one that supports the <, <=, >, and >= operators.
type Ordered interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
        ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
        ~float32 | ~float64 |
        ~string
}
```

##==##

# Orienté Objet - 10

## Les types génériques

<!-- .slide: class="with-code" -->

En plus des fonctions génériques, Go implémente aussi les types génériques.

Un type peut être paramétré, ce qui peut être intéressant pour implémenter des
structures génériques.

Voici l'exemple d'un noeud d'une liste chaînée générique en Go :

```Go
type LinkNode[T any] struct {
	Value T
	Next  *LinkNode[T]
}

func NewLinkNode[T any](value T) *LinkNode[T] {
	return &LinkNode[T]{
		Value: value,
		Next:  nil,
	}
}
```

##==##

# Orienté Objet - 10 - Exercice

Etant donné la structure de liste chaînée précédente,
ajouter les méthodes suivantes à la structure de donnée :

* Ajouter en tête
* Retourner un élément par son index (premier élément index `0`)

