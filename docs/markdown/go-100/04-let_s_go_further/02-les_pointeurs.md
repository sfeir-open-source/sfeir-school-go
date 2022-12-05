<!-- .slide: class="with-code" -->

# Pour aller plus loin - 03

## Les pointeurs

Go comporte des pointeurs. Un pointeur contient l'adresse mémoire d'une valeur.

Le type **`*T`** est un **pointeur** vers une valeur de type **T**.

```go
var p *int
```

<!-- .element: class="big-code" -->

Ici, p est un pointeur vers un entier.

##==##

<!-- .slide: class="with-code" -->

# Pour aller plus loin - 03

## Les pointeurs

La valeur zéro d’un pointeur est **nil**.

L'opérateur **&** génère un pointeur vers son opérande.

```Go
var p *int  // p vaut nil
i := 42
p = &i      // p pointe vers i
```

<!-- .element: class="big-code" -->

##==##

<!-- .slide: class="with-code" -->

# Pour aller plus loin -03

## Les pointeurs

L'opérateur \* dénote la valeur sous-jacente du pointeur.

```Go
i := 3
p := &i
fmt.Println(*p)   // écrit la valeur de i (=3) par le pointeur p
*p = 21           // définit la valeur de i à travers le pointeur p
```

<!-- .element: class="big-code" -->

##==##

<!-- .slide: class="with-code" -->

# Pour aller plus loin - 03

## Les pointeurs vers les structures

Les champs struct peuvent être accessibles via un pointeur de struct.
L'indirection via le pointeur est transparente.

```Go
func main() {
	v := Vector{1, 2}
	p := &v
	p.X = 7         // raccourci pour (*p).X = 7
	fmt.Println(v)  // { X:7, Y:2 }
}
```

<!-- .element: class="big-code" -->

##==##

<!-- .slide: class="first-slide" sfeir-level="2" sfeir-techno="Go" -->

Si vous appréciez la formation, Envoyez un Tweet !

#sfeirschool #golang

@Sfeirlille @sfeir

@sebastienfriess

##==##

<!-- .slide: class="first-slide" sfeir-level="2" sfeir-techno="Go" -->

Si vous appréciez la formation, Envoyez un Tweet !

#sfeirschool #golang

@sfeirstrbg @sfeir

@sebastienfriess

##==##

<!-- .slide: class="first-slide" sfeir-level="2" sfeir-techno="Go" -->

Si vous appréciez la formation, Envoyez un Tweet !

#sfeirschool #golang

@SFEIRLux @sfeir

@ogerardin

##==##

<!-- .slide:-->

# Pour aller plus loin - 03 - exercice

## Les pointeurs et les structures

Répondez aux questions en remplaçant REPLACEME par true ou false.
