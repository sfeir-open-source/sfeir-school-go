<!-- .slide: class="sfeir-bg-white-3" -->

# Les bases - 09

**Les boucles**

- Go a une seule structure de boucle, la boucle **for**.
- Elle a trois composantes séparées par des points-virgules:
- La déclaration d'initialisation : exécutée avant la première itération
- L'expression de condition : évaluée avant chaque itération
- La déclaration d'aboutissement : exécutée à la fin de chaque itération


##--##
<!-- .slide: class="with-code"-->

# Les bases - 09

**Les boucles**

Comme pour l’instruction **if**, il n'y a pas de parenthèses entourant les trois composantes de la déclaration **for** et les accolades { } sont exigées.

```Go
for i := 0; i < 10; i++ {
	sum += i
}
```
<!-- .element: class="big-code" -->


##--##
<!-- .slide: class="sfeir-bg-white-3" -->

# Les bases - 09

**Les boucles**

La déclaration d'initialisation sera souvent une déclaration de variable courte et les variables ainsi déclarées ne sont visibles que dans le cadre de la déclaration **for**.

L'itération de la boucle sera arrêtée une fois que la condition booléenne sera évaluée à false.


##--##
<!-- .slide: class="with-code" -->

# Les bases - 09

**Les boucles**

La déclaration d'initialisation et d'aboutissement sont facultatifs.

`
for ; n < 1000; {
	n += n
}
`

qui peut s’écrire sans ";"

`
for  n < 1000; {
	n += n
}
`


Notes:
=> équivalent du while



##--##
<!-- .slide: class="with-code" -->

# Les bases - 09

**Les boucles**

La condition peut elle aussi être omise.
On obtient ainsi une boucle infinie :

`for { }`


##--##
<!-- .slide: class="with-code" -->

# Les bases - 09 - Exercice

**Les boucles**

Écrire une fonction implémentant l'algorithme de la conjecture de Syracuse qui prend en paramètre :
- un entier de départ
- un nombre max d'itérations à effectuer

Et qui devra retourner :

- un booléen indiquant si le nombre 1 a été atteint
- le nombre d'itérations effectuées.


Notes:
si n est pair, n = n/2

si n est impair, n = n*3 + 1

=> la conjecture dit que quelque soit n de départ, on finit toujours par arriver à 1



##--##
<!-- .slide: class="with-code" -->

# Les bases - 10

**Switch**

Le switch fait exactement ce qu’on attend de lui. Comme pour le if ou le for, une expression peut être ajoutée avant la condition.
```Go
switch initialisation; expression {
    case a:
    case b:
    default:
}
```
<!-- .element: class="big-code" -->


##--##
<!-- .slide: class="sfeir-bg-white-3" -->

# Les bases - 10

**Switch**

En Go, un **case** se termine automatiquement (l’instruction **break** est inutile).
Si, au contraire, nous voulons que la **case** suivante soit exécutée, il faut ajouter le mot clé **fallthrough**.


##--##
<!-- .slide: class="with-code" -->

# Les bases - 10

**Switch**

Si l’on omet la condition, le switch peut permettre d'écrire élégamment de longues chaînes if-then-else.

```Go
switch {//équivalent à switch true
    case i < 0:
    case i == 0:
    default:
}
```
<!-- .element: class="big-code" -->
