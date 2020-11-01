
# Orienté Objet - 20

## Les méthodes

Go ne dispose pas de **classes**.

Cependant, on peut définir des méthodes sur les types.


##==##
<!-- .slide: class="with-code" -->
# Orienté Objet - 20

## Les méthodes

Une **méthode** est une **fonction** avec un **récepteur**, indiqué entre le mot clé **func** et le **nom** de la méthode.
```Go
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```
<!-- .element: class="big-code" -->

La méthode **Abs** possède un récepteur de type **Vertex** nommé **v**.


##==##
<!-- .slide: class="with-code" -->

# Orienté Objet - 20

## Les méthodes

Une **méthode** est donc juste une **fonction** avec un **récepteur**.

Voici Abs écrit comme une fonction régulière sans changement de fonctionnalité :

```Go
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```
<!-- .element: class="big-code" -->


##==##
<!-- .slide: class="with-code" -->

# Orienté Objet - 20

## Les méthodes

La différence se fait dans la façon d’appeler la fonction/méthode.

Méthode :

```
v := Vertex{1, 3}
a := v.Abs()
```
<!-- .element: class="big-code" -->

Fonction :

```
v := Vertex{1, 3}
a := Abs(v)
```
<!-- .element: class="big-code" -->

##==##
<!-- .slide: class="with-code" -->

# Orienté Objet - 20

## Les méthodes

On peut aussi déclarer une méthode sur des types non-struct.

````Go
type MyFloat float64
func (f MyFloat) IsPositive() bool {
	return f >= 0
}

````
<!-- .element: class="big-code" -->


##==##
<!-- .slide: class="with-code" -->

# Orienté Objet - 20

## Les méthodes

En revanche, le **type du récepteur** doit impérativement être déclaré dans le même **package** que la méthode.

Il est donc impossible d’écrire :

````Go
func (f float64) IsPositive() bool {
	return f >= 0
}
````
<!-- .element: class="big-code" -->

puisque float64 est un type défini dans un autre package.



##==##
<!-- .slide: class="with-code" -->

# Orienté Objet - 20

## Les méthodes, récepteur de pointeur

On peut déclarer des méthodes avec des récepteurs de **pointeur**.

Cela permet à la méthode de modifier la valeur vers laquelle le récepteur pointe :

````GO
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
````
<!-- .element: class="big-code" -->



Notes:
Avec un récepteur de valeur, la méthode Scale fonctionnerait sur une copie de la valeur originale Vertex, et n’aurait donc pas l’effet escompté.


##==##
<!-- .slide: class="with-code" -->

# Orienté Objet - 21

## Les méthodes, récepteur de pointeur

Un autre avantage des récepteurs de pointeurs est de pouvoir profiter de la transparence de l’indirection de Go :

`v.Scale(5)`  compilera si **v** est de type ***Vertex** ou **Vertex**.

alors que :

Si on a `func ScaleFunc(v *Vertex, f float64)`

**ScaleFunc(v, 10)** ne compilera que si **v** est de type ***Vertex**.


##==##

# Orienté Objet - 21

Choix d'un récepteur de valeur ou pointeur

Il y a deux raisons d'utiliser un récepteur de pointeur :

- le cas ou la méthode doit modifier la valeur du récepteur pointé.

- éviter la copie de la valeur sur chaque appel de méthode.

⇒ à l’usage on utilisera le plus souvent des récepteurs de pointeur.


##==##
<!-- .slide: class="with-code" -->

# Orienté Objet - 22

## Le récepteur nil

En Go, il est possible d’appeler une méthode sur un pointeur **nil** sans générer d’exception :

```Go
var p *T = nil
p.SomeMethod()
```
<!-- .element: class="big-code" -->


##==##
<!-- .slide: class="with-code" -->

# Orienté Objet - 22

## Le récepteur nil

Il faut malgré tout que la méthode soit prévue pour fonctionner avec un récepteur ayant la valeur **nil** :

````Go
func (pf *MyFloat) IsPositive() bool {
	return pf != nil && *pf >= 0
}
var p *MyFloat = nil
ok := p.IsPositive() // pas d’erreur. ok = false
````
<!-- .element: class="big-code" -->



##==##

# Orienté Objet - 22 - exercice

## Les méthodes et le récepteur nil

Créez un type **person** qui contiendra le prénom et le nom d'une personne.

Ajoutez une méthode **fullname()** sur le type ***person** qui retourne la concaténation du prénom et du nom.
La méthode devra gérer le cas où le récepteur est nil.





