# Orienté Objet - 10

## Les Fonctions Génériques

Les fonctions génriques permettent de manipuler plusieurs type de paramètres.

```golang
func MyFunc[T any | int64 | float64](collection []T) T {}
```

##==##

<!-- .slide: class="with-code" -->

# Orienté Objet - 10

## Les Fonctions Génériques

Exemple sans Génériques:

```golang

func FilterInterface(collection []interface{}, predicate func(iten interface{}, index int) bool) []interface{} {
 result := make([]interface{}, 0)
 for i, item := range collection {
  if predicate(item, i) {
   result = append(result, item)
  }
 }

 return result

}

func main() {

 listInterface := append(make([]interface{}, 0), 1, "2", "3", 4)
 resultInterface := FilterInterface(listInterface, func(data interface{}, index int) bool {
  if nbr, ok := data.(int); ok != false {
   return nbr%2 == 0
  }
  return false
 })
 fmt.Printf("%v", resultInterface)

```

<!-- .slide: class="with-code" -->

# Orienté Objet - 10 - exercices
## Les Fonctions Génériques

Le même exemple avec les Génériques:

```golang

func Filter[V comparable](collection []V, predicate func(item V, index int) bool) []V {
 result := []V{}

 for i, item := range collection {
  if predicate(item, i) {
   result = append(result, item)
  }
 }

 return result
}

type MyStruct struct {
 myNbr int
}

func main() {

 list := []MyStruct{{myNbr: 1}, {myNbr: 2}, {myNbr: 3}, {myNbr: 4}}

 result := Filter(list, func(nbr MyStruct, index int) bool {
  return nbr.myNbr%2 == 0
 })
 fmt.Printf("%v", result)

 listInt := []int{1, 2, 3, 4}

 resultInt := Filter(listInt, func(nbr int, index int) bool {
  return nbr%2 == 0
 })

 fmt.Printf("%v", resultInt)
}
```
