<!-- .slide: class="sfeir-bg-white-3" -->

# Pour aller plus loin - 14


Go comporte des pointeurs. Un pointeur contient l'adresse mémoire d'une valeur.
Le type *T est un pointeur vers une valeur de type T.
var p *int
Ici, p est un pointeur vers un entier.
Les pointeurs


##==##
<!-- .slide:-->

# Pour aller plus loin - 14


La valeur zéro d’un pointeur est nil.
L'opérateur & génère un pointeur vers son opérande.
var p *int  // p vaut nili := 42p = &i  // p est un pointeur vers i



Les pointeurs


##==##
<!-- .slide:-->

# Pour aller plus loin - 14


L'opérateur * dénote la valeur sous-jacente du pointeur.
i := 3p := &ifmt.Println(*p)   // lire i (=3) par le pointeur p*p = 21                // définir i par le pointeur p



Les pointeurs


##==##
<!-- .slide:-->

# Pour aller plus loin - 14


Les champs struct peuvent être accessibles via un pointeur de struct.
L'indirection via le pointeur est transparente.
func main() {	v := Vector{1, 2}	p := &v	p.X = 7	fmt.Println(v)  // { X:7, Y:2 }}


Les pointeurs vers les structures


##==##
<!-- .slide:-->




GO


Si vous appréciez la formation, Envoyez un Tweet !


#sfeirschool #golang
@Sfeirlille @sfeir
@sebastienfriess


![](./images/school_Badge_blanc.png)

![](./images/STAR.png)

![](./images/STAR.png)

##==##
<!-- .slide:-->




GO


Si vous appréciez la formation, Envoyez un Tweet !


#sfeirschool #golang
@sfeirstrbg @sfeir
@sebastienfriess


![](./images/school_Badge_blanc.png)

![](./images/STAR.png)

![](./images/STAR.png)

##==##
<!-- .slide:-->




GO


Si vous appréciez la formation, Envoyez un Tweet !


#sfeirschool #golang
@SFEIRLux @sfeir
@ogerardin


![](./images/school_Badge_blanc.png)

![](./images/STAR.png)

![](./images/STAR.png)

##==##
<!-- .slide:-->

# Pour aller plus loin - 14 - exercice


Répondez aux questions en remplaçant REPLACEME par true ou false.


Les pointeurs et les structures

