<!-- .slide: class="with-code" -->

# Pour aller plus loin - 12

## Déclaration de type

Nous pouvons déclarer un nouveau type à l’aide du mot clé **type**.

Par exemple, il est possible de déclarer un type dont le type sous-jacent est un nombre :

`type Distance int64`

C’est exactement ce qui est fait dans le package time de Go :

`type Duration int64`


##--##
<!-- .slide: class="with-code" -->

# Pour aller plus loin - 12

## Déclaration de type

Il est possible d’effectuer les même opérations que sur le type sous-jacent, mais seulement entre variables du même type. C’est très utile pour éviter les erreurs d’homogénéité dans une formule.

```Go
var d Distance = 300
var t Duration = 60
v := d / t // erreur de compilation
v := float64(d) / float64(t)  // ok
```
<!-- .element: class="big-code" -->

Notes:
Il est possible de repasser au type sous-jacent par exemple int64(d)


##--##
<!-- .slide: class="with-code" -->
# Pour aller plus loin - 13

## Les structures

Une structure est une collection de champs. Déclarer une structure, c’est déclarer un nouveau type. Il est donc logique de déclarer une structure à l’aide du mot clé type et du type sous-jacent struct :
```Go
type Vector struct {
	X int
	Y int
}
```
<!-- .element: class="big-code" -->

##--##
<!-- .slide: class="with-code" -->

# Pour aller plus loin - 13

## Les structures

Une structure peut être instanciée littéralement, offrant la possibilité d’initialiser la valeur des champs :
```Go
v1 := Vector{1, 2}  // X:1 et Y:2
v2 := Vector{X: 1}  // X:1, et Y:0 est implicite
v3 := Vector{}        // X:0 et Y:0
```
<!-- .element: class="big-code" -->


##--##
<!-- .slide: class="with-code" -->

# Pour aller plus loin - 13

## Les structures

Les champs de la structure sont accessibles à l'aide d'un point :
```Go
func main() {
	v := Vector{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
```
<!-- .element: class="big-code" -->
