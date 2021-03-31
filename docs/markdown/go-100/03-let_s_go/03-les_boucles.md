# Les bases - 09

## Les boucles

- Go a une seule structure de boucle, la boucle **for**.
- Elle a **trois composantes** séparées par des points-virgules:
  - L'**initialisation** : exécutée avant la première itération
  - La **condition** de continuation : évaluée avant chaque itération, si elle devient fausse on sort.
  - L'instruction de **pré-bouclage** : exécutée à la fin de chaque itération

##==##

<!-- .slide: class="with-code"-->

# Les bases - 09

## Les boucles

Comme pour l’instruction **if**, il n'y a pas de parenthèses entourant les trois composantes de la déclaration **for** et les accolades { } sont exigées.

```Go
for i := 0; i < 10; i++ {
	sum += i
}
```

<!-- .element: class="big-code" -->

##==##

# Les bases - 09

## Les boucles

La déclaration d'initialisation sera souvent une déclaration de variable courte et les variables ainsi déclarées ne sont visibles que dans le cadre de la déclaration **for**.

L'itération de la boucle sera arrêtée une fois que la condition de continuation sera évaluée à _false_.

##==##

<!-- .slide: class="with-code" -->

# Les bases - 09

## Les boucles

L'initialisation et l'instruction de pré-bouclage sont facultatives.

```go
n :=1
for ; n < 1000; { n += n }
```

<!-- .element: class="big-code" -->

qui peut s’écrire sans ";"

```go
n := 1
for n < 1000 { n += n }
```

<!-- .element: class="big-code" -->

Notes:
=> équivalent du while

##==##

<!-- .slide: class="with-code" -->

# Les bases - 09

## Les boucles

La condition peut elle aussi être omise.
On obtient ainsi une boucle infinie :

```go
for { … }
```

dont on peux heureusement sortir ! (instruction _break_)

<!-- .element: class="big-code" -->

##==##

<!-- .slide: class="with-code" -->

# Les bases - 09 - Exercice

## Les boucles

Écrire une fonction implémentant l'algorithme de la conjecture de Syracuse qui prend en paramètre :

- un entier de départ
- un nombre max d'itérations à effectuer

Et qui devra retourner :

- un booléen indiquant si le nombre 1 a été atteint
- le nombre d'itérations effectuées.

Dans la suite de syracuse le suivant _S_ est, en fonction du précedent _P_ :

- _P_/2 si _P_ est pair
- 3\*_P_+1 si _P_ est impair

Le but du jeu est, de suivant en suivant, d'atteindre le nombre 1, sans dépasser un nombre maximal d'itérations fixés à l'avance.

Notes:
Aussi appellée conjecture de Collatz.
=> la conjecture dit que quelque soit n de départ, on finit toujours par arriver à 1

C'est une conjecture, c'est à dire que ce n'est pas démontré mathématiquement, même si toujours testé positif jusqu'à présent.
##==##

<!-- .slide: class="with-code" -->

# Les bases - 10

## Switch

Le **switch** fait exactement ce qu’on attend de lui. Comme pour le **if** ou le **for**, une expression peut être ajoutée avant la condition.

```Go
switch initialisation; expression {
    case a:
    case b:
    default:
}
```

<!-- .element: class="big-code" -->

##==##

# Les bases - 10

## Switch

En Go, un **case** se termine automatiquement (l’instruction **break** est inutile).
Si, au contraire, nous voulons que la **case** suivante soit exécutée, il faut ajouter le mot clé **fallthrough**.

##==##

<!-- .slide: class="with-code" -->

# Les bases - 10

## Switch

Si l’on omet la condition, le switch peut permettre d'écrire élégamment de longues chaînes if-then-else.

```Go
switch { // équivalent à switch true
    case i < 0:
    case i == 0:
    default:
}
```

<!-- .element: class="big-code" -->
