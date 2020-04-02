<!-- .slide: class="sfeir-bg-white-3" -->

# Orienté Objet - 24

## Les interfaces

En Go, il n’y a pas de notion d’héritage entre les **struct**, mais la notion d’**interface** est bien présente.


##--##
<!-- .slide: class="with-code" -->

# Orienté Objet - 24

## Les interfaces

Comme on s’y attend, un type d'interface est définie comme un ensemble de signatures de méthode.

````Go
type Oiseau interface {
	Vole(direction Vector)
	Mange() (int, bool)
}
````
<!-- .element: class="big-code" -->


##--##
<!-- .slide: class="sfeir-bg-white-3" -->

# Orienté Objet - 25

## Les interfaces

Particularité de Go : les interfaces sont implémentées **implicitement**.

Un type implémente “automatiquement” une **interface** en mettant en œuvre ses méthodes.

⇒ Il n'y a aucune déclaration explicite d'intention, pas de mot-clé «implements».



Notes:
Les interfaces implicites découple la définition d'une interface de son implémentation qui pourraient alors apparaître dans n'importe quel paquet sans arrangement préalable.



##--##
<!-- .slide: class="with-code" -->
# Orienté Objet - 25

## Les interfaces

Si on a :

````Go
type I interface {
	M()
}
type T struct { }
````
<!-- .element: class="big-code" -->


Alors il suffit de déclarer la méthode :

````Go
func (t T) M() {
	/*...*/
}
````
<!-- .element: class="big-code" -->

pour que le type T implémente I.


##--##
<!-- .slide: class="with-code" -->

# Orienté Objet - 25

## Les interfaces

Si on a :
```Go
type Aigle struct {}
func (a Aigle) Mange() { /*...*/ }
func (a Aigle) Vole() { /*...*/ }
```
<!-- .element: class="big-code" -->

Et :
```Go
type Oiseau interface {
	Mange()
	Vole()
}
type Volant interface { Vole() }
```
<!-- .element: class="big-code" -->


⇒ Quelle(s) interface(s) Aigle implémente-t-il ?


##--##
<!-- .slide: class="with-code" -->

# Orienté Objet - 25

## Les interfaces

Si on a :

```Go
type Aigle struct {}
func (a Aigle) Mange() { /*...*/ }
func (a Aigle) Vole() { /*...*/ }
```
<!-- .element: class="big-code" -->

Et :

```Go
type Oiseau interface {
	Mange()
	Vole()
}
type Volant interface { Vole() }
```
<!-- .element: class="big-code" -->

⇒ **Aigle** implémente à la fois l’interface **Oiseau** et l’interface **Volant**.


##--##
<!-- .slide: class="with-code" -->

# Orienté Objet - 25

## Les interfaces

En revanche, si on a :

```Go
type A380 struct {}
func (a A380) Vole() { /*...*/ }
```
<!-- .element: class="big-code" -->


Et :

```Go
type Oiseau interface {
	Mange()
	Vole()
}
type Volant interface { Vole() }
```
<!-- .element: class="big-code" -->


⇒ Quelle(s) interface(s) A380 implémente-t-il ?



##--##
<!-- .slide: class="with-code" -->

# Orienté Objet - 25

## Les interfaces

En revanche, si on a :

```Go
type A380 struct {}
func (a A380) Vole() { /*...*/ }
```
<!-- .element: class="big-code" -->

Et :

```Go
type Oiseau interface {
	Mange()
	Vole()
}
type Volant interface { Vole() }
```
<!-- .element: class="big-code" -->

⇒ **A380** implémente seulement l’interface **Volant**.


##--##
<!-- .slide: class="with-code" -->

# Orienté Objet - 26

## Les interfaces

Attention,

il suffit de déclarer la méthode :

```Go
func (t T) M() {
	fmt.Println(t.S)
}
```
<!-- .element: class="big-code" -->

pour que le type **T** implémente **I**.


mais si on déclare :

```Go
func (t *T) M() {
	fmt.Println(t.S)
}
```
<!-- .element: class="big-code" -->

c’est le type ***T** qui implémente **I**.


##--##
<!-- .slide: class="sfeir-bg-white-3" -->

# Orienté Objet - 26

## Les interfaces

⇒ En général, toutes les méthodes sur un type donné **T** doivent avoir soit un récepteur de valeur (T) soit un récepteur de pointeur (*T), mais pas un mélange des deux.

Sinon, **T** implémenterait des interfaces tandis que ***T** en implémenterait d’autres, ce qui deviendrait pénible à utiliser.



##--##
<!-- .slide: class="sfeir-bg-white-3" -->

# Orienté Objet - 27

## L’interface vide

L’interface vide ne définit, comme son nom l’indique, aucune méthode :

**interface{}**

Tous les types implémentent l’interface vide puisqu’un type a toujours au moins zéro méthode. On utilise donc l’interface vide pour gérer des valeurs de type inconnu.

Ex: *fmt*.Print prend des arguments de type **interface{}**


##--##
<!-- .slide: class="with-code" -->

# Orienté Objet - 28

## Assertion de type

Une assertion de type donne accès à la valeur concrète **(sous-jacente)** d'une valeur d'interface.

`t := i.(T)`

Cette instruction affirme que la valeur d'interface **i** contient le type concret **T** et assigne la valeur sous-jacente à la variable **t**.



##--##
<!-- .slide: class="with-code" -->

# Orienté Objet - 28

## Assertion de type

Si **i** ne détient pas **T**, la déclaration va déclencher une **panique**.

Mais l’assertion de type peut aussi s’effectuer sous la forme :

`t, ok := i.(T)`

Si i ne détient pas **T**, aucune panique ne sera déclenchée, mais **ok** aura la valeur **false**.


##--##
<!-- .slide: class="with-code" -->

# Orienté Objet - 29

## Switch de type

Un switch de type permet de gérer les différents types concrets qu’une interface pourrait avoir :

```Go
switch v := i.(type) {
case T:    // ici v est de type T
case S:   // ici v est de type S
default: // ici v a le même type que i
}
```
<!-- .element: class="big-code" -->


##--##
<!-- .slide: class="sfeir-bg-white-3" -->

# Orienté Objet - 29 - exercice

## Les interfaces

1. Créer un type **FmtLogger** qui implémente l’interface **Logger**.

2. Finir l’implémentation de la fonction **recoverPanic** en utilisant le **logger** pour écrire ce que contient la panique **r**.

3. Instancier le **FmtLogger** et appeler la fonction **recoverPanic** dans la fonction **main**.


