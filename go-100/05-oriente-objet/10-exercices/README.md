# Les génériques

Etant donné une structure de type liste chaînée suivante :

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

type LinkedList[T any] struct {
    Head   *LinkNode[T]
    length int
}

func NewLinkedList[T any]() *LinkedList[T] {
    return &LinkedList[T]{
            Head:   nil,
            length: 0,
	  }
}
```

Ajouter les méthodes suivantes à cette structure de donnée :

* Ajouter en tête
* Retourner un élément par son index (premier élément index `0`)
