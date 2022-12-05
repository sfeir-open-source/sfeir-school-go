<!-- .slide: class="with-code" -->

# Pour aller plus loin - 04

## Les tableaux

Le type **[n]T** est un tableau de n valeurs de type **T**.

L’expression :

```go
var a [10]int
```

<!-- .element: class="big-code" -->

déclare une variable **a** comme un tableau de dix entiers.

##==##

# Pour aller plus loin - 04

## Les tableaux

La longueur d'un tableau fait partie de son type.

⇒ les tableaux ne peuvent pas être redimensionnés.

Notes:
La longueur d'un tableau fait partie de son type, de sorte que les tableaux ne peuvent pas être redimensionnés. Cela semble limitatif, mais ne vous inquiétez pas ; Go offre un moyen pratique de travailler avec des tableaux.

##==##

# Pour aller plus loin - 05

## Les slices

Un **slice** pointe vers un tableau de valeurs et possède une longueur.

**[]T** est un slice avec des éléments de type **T**.

**len(s)** retourne la longueur du slice **s**.

<img class="float-right" src="./assets/go-100/images/slice-1-array.png">

cap(s) retourne la capacité du slice **s**.

La valeur zéro d’un slice est **nil**.

##==##

<!-- .slide: class="with-code" -->

# Pour aller plus loin - 05

## Les slices

Les slices peuvent être créés avec la fonction **make**.

Cela alloue un tableau vide et retourne un slice qui référence ce tableau :

```go
a := make([]int, 5) // len(a)=5
```

<!-- .element: class="big-code" -->

##==##

# Pour aller plus loin - 05

## Les slices

Les slices peuvent être redécoupés, créant ainsi un nouveau slice qui pointe vers le même tableau.

L’expression :

**s[low:high]**

récupère les éléments du slice de **low** à **high-1** inclus (ou **high** exclus).

##==##

<!-- .slide: class="two-column-layout with-code" -->

# Pour aller plus loin - 05

## Les slices

##--##

<br><br>

<img src="./assets/go-100/images/slice-2-array.png">

<br>

<img src="./assets/go-100/images/slice-2-array-decoupe.png">

##--##

<!-- .slide: class="with-code" -->

<br><br>

```Go
var a = make([]byte, 5)
len(a) == 5
cap(a) == 5
```

<!-- .element: class="big-code" -->

<br>

```Go
var b = a[2:4]
len(b) == 2
cap(b) == 3
```

<!-- .element: class="big-code" -->

##==##

<!-- .slide: class="with-code" -->

# Pour aller plus loin - 05

## Les slices

Pour spécifier une capacité, passer un troisième argument à make :

```Go
b := make([]int, 0, 5)      // len(b)=0, cap(b)=5
b = b[:cap(b)]             // len(b)=5, cap(b)=5
b = b[1:]                 // len(b)=4, cap(b)=4
```

<!-- .element: class="big-code" -->

##==##

<!-- .slide: class="with-code" -->

# Pour aller plus loin - 05

## Les slices

Il est courant d'ajouter de nouveaux éléments à un slice, et ainsi Go offre une fonction intégrée **append**.

`func append(s []T, vs ...T) []T`

Le premier paramètre s de append est un slice de type T, et les suivants sont des valeurs T à ajouter au slice.

Notes:
La valeur résultante de append est une tranche contenant tous les éléments de la tranche original et les valeurs fournis.

Si le tableau de support de s est trop petit pour contenir toutes les valeurs données un plus gros tableau sera alloué. La tranche retourné pointera vers le tableau nouvellement allouée.

##==##

# Pour aller plus loin - 06

## For … range

La forme **range** de la boucle **for** permet d’itérer sur tous les éléments d’un slice, d’un tableau, ou d’une map.

Deux valeurs sont renvoyées pour chaque itération :
l'indice et une **copie** de l'élément à cet indice.

##==##

<!-- .slide: class="with-code" -->

# Pour aller plus loin - 06

## For … range

```go
for i, v := range s { … }
```

<!-- .element: class="big-code" -->

Correspond à :

```go
for i := 0; i < len(s); i++ {
  v := &s[i]
}
```

<!-- .element: class="big-code" -->

mais ici \*v est modifiable

##==##

<!-- .slide: class="with-code" -->

# Pour aller plus loin - 06

## For … range

On peut omettre l’indice en le remplaçant par “\_” :

```go
for _, v := range s { … }
```

<!-- .element: class="big-code" -->

On peut aussi omettre la valeur.

```go
for i := range s { … }
```

<!-- .element: class="big-code" -->

##==##

# Pour aller plus loin - 06 - exercice

## Les slices et le range

Écrire les fonctions **indexOf** et **split**.

Notes:
-> 2h

##==##

<!-- .slide: class="with-code" -->

# Pour aller plus loin - 07

## Maps

Une map est une liste de clés-valeurs. Une map se déclare :

`map[typeClé]typeValeur`

Les maps doivent être créées avec make. Leur valeur-zéro est nil.

```go
m = make(map[string]Vector)
```

<!-- .element: class="big-code" -->

##==##

<!-- .slide: class="with-code" -->

# Pour aller plus loin - 07

## Maps

Comme les structs, une map peut être instanciée littéralement.

Il faut cependant indiquer la valeur de chaque clé :

```Go
var m = map[string]Vector{
	"Bell Labs": Vector{40.68433, -74.39967},
	"Google": Vector{37.42202, -122.08408},
}
```

<!-- .element: class="big-code" -->

##==##

<!-- .slide: class="with-code" -->

# Pour aller plus loin - 07

## Maps

- Affectation : `m[key] = elem`
- Suppression : `delete(m, key)`
- Récupération :

```go
  elem = m[key]     // elem prend la valeur-zéro de son type si la clé n'est pas présente.

  elem, ok = m[key] // ok est true si la clé est présente, sinon false.
```

<!-- .element: class="big-code" -->

Notes:
Insérer ou mettre à jour un élément de carte m :

m[key] = elem

Récupérer un élément :

elem = m[key]

Supprimer un élément :

delete(m, key)

Test qu'une clé est présente avec une affectation de deux valeurs :

elem, ok = m[key]

Si key est dans m , ok est true. Sinon, ok est false

Si key n'est pas dans la carte, alors elem est la valeur zéro pour le type d'élément de la carte.

Note :
si elem ou ok
ont pas encore été déclarée, vous pouvez utiliser une déclaration courte :

elem, ok := m[key]

##==##

<!-- .slide: class="with-code" -->

# Pour aller plus loin - 08

## Un peu plus sur les fonctions

En Go, les fonctions sont aussi des valeurs, elles peuvent être utilisées comme arguments de fonctions ou valeurs de retour.

On dit que les fonctions sont **first class citizen** en Go.

```go
type F func(a int) int

func compose(g, h F) F {
	return func(n int) int {
		return g(h(n))
	}
}

var square, increment, incThenSquare F

square = func(n int) int { return n * n }
increment = func(n int) int { return n + 1 }

incThenSquare = compose(square, increment)
```

##==##

# Pour aller plus loin - 08

## Closures

Les fonctions de Go peuvent être des **closures**.

Une **closure** (fermeture) est une valeur de la fonction qui fait référence à des variables à partir de l'extérieur de son corps.

La fonction peut accéder et assigner les variables référencées, dans ce sens, la fonction est «liée» aux variables.

##==##

<!-- .slide: class="with-code" -->

# Pour aller plus loin - 08 - exercice

## Closures

Implémentez la fonction **multiplicateurPar** pour qu'elle retourne une fonction qui multiplie son argument par un nombre donné.

```Go
// fonction qui retourne une fonction qui retourne un int.
func multiplicateurPar(x int) func(int) int { }
```

<!-- .element: class="big-code" -->
