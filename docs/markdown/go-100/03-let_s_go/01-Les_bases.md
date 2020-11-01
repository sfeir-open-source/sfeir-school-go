<!-- .slide: class="with-code" -->

# Les bases - 01

## Hello, 世界

```Go
package main

import "fmt" //Le programme est organisé en packages.
             //Il doit avoir à minima un package main, contenant une fonction main.

func main() { //Go fournit des centaines de packages (fmt, net, http, json, etc.)
   fmt.Println("Hello, 世界") //La fonction main est le point d’entrée du programme.
}
```
<!-- .element: class="big-code" -->

Notes:
Demander aux élèves d’ouvrir le fichier 01/main.go dans VSC.
Go supporte nativement l’UTF8. Tous les fichiers sources doivent être encodés en UTF8.



##==##

# Les bases - 01

## Hello, 世界

Remarque : comme on s’y attend, 世界 signifie “monde”, mais ces caractères ne sont pas disponibles en ANSI.

- Par convention, en Go, tous les fichiers sources sont encodés en **UTF8**.
- Les chaînes de caractères sont encodées en **UTF8** par défaut.
- Il est possible de nommer une variable **Δt** par exemple.

Notes:
On peut aussi nommer une variable Δt par exemple.



##==##
<!-- .slide: class="with-code" -->

# Les bases - 01

## Compilation, exécution

Placez vous dans le répertoire 01, puis :
- Compilez le programme : `$ go build main.go`
- Ensuite, exécutez le programme : `$ ./main`
- Ou plus simplement : `$ go run main.go`

Go est un langage compilé. Le compilateur produit un binaire exécutable.
Pour compiler pour un autre OS :
` $ GOOS=linux GOARCH=amd64 go build main.go`


Notes:
Heureusement un même code Go peut être compilé pour Linux, MacOs, Windows, BSD, Unix, et même les smartphones.

GOOS=linux GOARCH=amd64



##==##

# Les bases - 02

## Imports

- En Go, il est **interdit** d’importer un package sans l’utiliser.
- Pour importer plusieurs packages, on préférera la notation factorisée.
- Par convention, le nom du package est le même que le dernier élément du chemin d'importation.Ex: chemin “math/rand” ⇒ package “rand”


##==##

# Les bases - 03

## Exporter un identifiant

Pour qu’un identifiant (de variable, fonction, type, etc.) soit visible en dehors du package dans lequel il est déclaré, il doit commencer par une majuscule. Il n’y a pas de mot clé “public”, “private”, ou autre en Go.

**⇒ un identifiant est exporté (ie. public) si, et seulement si, il commence par une majuscule.**


##==##
<!-- .slide: class="with-code" -->

# Les bases - 04

## Les fonctions

Une fonction est déclarée à l’aide du mot clé **func** suivi du nom de la fonction. Le corps de la fonction est compris entre { et }. Le caractère ‘{‘ doit être obligatoirement sur la même ligne que le mot clé **func**.

`
    func maFonction() { }
`


##==##
<!-- .slide: class="with-code" -->

# Les bases - 04

## Les fonctions

Une fonction peut prendre zéro ou plusieurs arguments.

`func maFonction(n int, s string) { }`

Notez que le type vient **après** l'identifiant de la variable.


##==##
<!-- .slide: class="with-code" -->

# Les bases - 04

## Les fonctions

Une fonction peut retourner une valeur.

```Go
func maFonction(n int, s string) string {
    return s + strconv.Itoa(n)
}
```
<!-- .element: class="big-code" -->



##==##
<!-- .slide: class="with-code" -->

# Les bases - 04

## Les fonctions

En Go, une fonction peut aussi retourner **plusieurs** valeurs.

```Go
func maFonction(n int, s string) (string, bool) {
	return s + strconv.Itoa(n), n == 0
}
```
<!-- .element: class="big-code" -->



##==##
<!-- .slide: class="with-code" -->

# Les bases - 04

## Les fonctions

Les valeurs de retour peuvent aussi être nommées.


```Go
func maFonction(n int, s string) (result string, isZero bool) {
	result = s + strconv.Itoa(n)
	isZero = n == 0
	return
}

```
<!-- .element: class="big-code" -->


Notes:
Les valeurs de retour de Go peuvent être nommés. Si c'est le cas, elles sont traités comme des variables définies dans le haut de la fonction.

Ces noms doivent être utilisés pour documenter la signification des valeurs de retour.

Une déclaration
return
 sans arguments renvoie les valeurs de retour nommés. C'est ce qu'on appelle un retour « nu ».

Des déclarations de retour nus doivent être utilisés uniquement dans des fonctions courte, comme avec l'exemple montré ici. Ils peuvent nuire à la lisibilité de fonctions plus grande.


##==##
<!-- .slide: class="with-code" -->

# Les bases - 04

## Les fonctions

Lorsque deux ou plusieurs paramètres de fonction consécutifs partagent un type, vous pouvez omettre le type de tous sauf le dernier.

```Go
func maFonction(a bool, b int, c int, d int, e string, f string) { ... }
⇔
func maFonction(a bool, b, c, d int, e, f string) { ... }
```
<!-- .element: class="big-code" -->

##==##
<!-- .slide: class="with-code" -->

# Les bases - 05

## Les variables

On déclare une (ou plusieurs) variable avec l’instruction **var** suivi du nom de la variable, puis du type et de sa valeur initiale.

`var nom type = expression`

Le type peut être omis si on donne une valeur initiale, et vice versa.


##==##
<!-- .slide: class="with-code" -->

# Les bases - 05

## Les variables

Concrètement, voici trois façons de déclarer une même variable :

```Go
var nom string = ""    //déclaration inutilement complète

var nom = ""           //le type est inféré par la valeur initiale (la chaîne vide)

var nom string         //la valeur initiale est la valeur-zéro du type
```
<!-- .element: class="big-code" -->



Notes:
La valeur-zéro du type string est la chaîne vide.

Les déclarations de variables peuvent être « factorisé », en blocs, comme les imports.


##==##
<!-- .slide: class="with-code" -->

# Les bases - 05

## Les variables

Il est possible de factoriser la déclaration de plusieurs variables à l’aide de parenthèses :

```Go
var (
    toto = ""
    titi float32
    tata = true
)
```
<!-- .element: class="big-code" -->


##==##
<!-- .slide: class="with-code" -->

# Les bases - 05

## Les variables

On peut aussi déclarer plusieurs variables sur une même ligne :

```Go
var a, b, c = "Hello", 42, true
```
<!-- .element: class="big-code" -->

Ou si elles sont du même type :
```Go
var x, y, z float64
```
<!-- .element: class="big-code" -->


##==##
<!-- .slide: class="with-code" -->

# Les bases - 05

## Les variables

A l’intérieur d’une fonction, on peut utiliser la déclaration *courte* :

```Go
func plop() {
	a := "plop!" // équivalent à : var a = "plop!"
	fmt.Println(a)
}
```
<!-- .element: class="big-code" -->


##==##

# Les bases - 05

## Les variables

Attention, gardez à l’esprit que <span style="color:blue">:=</span> est une <span style="color:blue">déclaration+assignation</span>, tandis que <span style="color:green">=</span> est une <span style="color:green">assignation</span>

Pour cette raison, il est <span style="color:red">impossible</span> d’assigner une valeur à une variable <span style="color:red">déjà déclarée</span> en utilisant <span style="color:red">:=</span>


##==##

# Les bases - 05

## Les variables

**A savoir :** comme pour la déclaration, on peut assigner plusieurs variables sur une même ligne :

`x, y := 42.0, 3.14`

Très pratique pour faire un _swap_ :

`x, y = y, x`


##==##

# Les bases - 06

## Les types de base

- bool
- string
- int, int8, int16, int32, int64
- uint, uint8, uint16, uint32, uint64

**int** est soit un **int32**, soit un **int64**, dépendamment de la plateforme et du compilateur. Il en va de même pour **uint** qui est soit **uint32** soit **uint64**.


##==##

# Les bases - 06

## Les types de base

- **byte** qui est un alias de int8
- **rune** qui est un alias de int32 et représente un “code point” Unicode
- **uintptr** qui est un entier non signé assez grand pour contenir la valeur d’un pointeur (généralement 32 ou 64 bits).


##==##

# Les bases - 06

## Les types de base

- **float32, float64**
- **complex64, complex128**

Le type **complex64** est un nombre complexe dont la partie réelle et la partie imaginaire sont des **float32**. Le type **complexe128** utilise des **float64**.


##==##

# Les bases - 06

## Les valeurs zéro

Chaque type a une **valeur zéro**. Une variable déclarée sans valeur initiale (ex: var nom type) aura pour valeur la **valeur zéro** de son type.
Les valeurs zéro des types de base sont :
- 0 pour les types numériques,
- false pour les types booléen, et
- "" (La chaîne vide) pour les chaînes.

##==##
<!-- .slide: class="with-code" -->
# Les bases - 06

## Les conversions de type

En Go, toute conversion de type doit être **explicite**.

L'expression T(v) convertit la valeur v au type T.

````Go
var i int = 42
var f float64 = float64(i)
````
<!-- .element: class="big-code" -->


##==##

# Les bases - 07

## Les constantes

- Les constantes sont déclarées comme des variables, mais avec le mot-clé **const**.

- Les constantes peuvent être un caractère, une chaîne, un booléen, ou des valeurs numériques.

- Les constantes ne peuvent pas être déclarées avec la syntaxe := .


##==##
<!-- .slide: class="with-code" -->

# Les bases - 07

## Les constantes

A savoir : les constantes numériques sont des valeurs de haute précision.

```Go
const (
	Big   = 1 << 100
		Small = Big >> 99
    )
```
<!-- .element: class="big-code" -->


Notes:
Les déclarations de constantes peuvent être « factorisé », en blocs, comme les imports et les déclarations de variables.


##==##
<!-- .slide: class="with-code" -->

# Les bases - 07

## Les constantes

A savoir : mot clé **<span style="color: green">iota</span>**

```Go
const (
	Small = iota	// == 0
	Medium     	// == 1
    Big			// == 2
)
```
<!-- .element: class="big-code" -->


Notes:
Les déclarations de constantes peuvent être « factorisé », en blocs, comme les imports et les déclarations de variables.
