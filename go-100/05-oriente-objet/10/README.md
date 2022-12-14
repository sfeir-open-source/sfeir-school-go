# Les fonctions génériques

Une première version des génériques est déjà possible grâces aux `interface{}`.

Mais les types et fonctions génériques vont un cran plus loin.

```golang
func Index[T comparable](s []T, x T) int
```

Les fonctions génériques permettent de fonctionner avec plusieurs type de paramètres.

Les types apparaissent `[]` avant la déclaration des paramètres de la fonction.

Le type `comparable` est une contrainte qui permet d'utiliser les opérateur `==` et `!=`
sur les valeurs des types.

# Les types génériques

En plus des fonctions génériques, Go implémente aussi les types génériques.

Un type peut être paramétré, ce qui peut être intéressant pour implémenter des
structures génériques.

Voici l'exemple d'un noeud d'une liste chaînée générique en Go :

```golang
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
