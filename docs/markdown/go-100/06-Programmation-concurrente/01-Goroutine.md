<!-- .slide: class="with-code" -->

# Concurrence - 29

## Goroutine

Une **goroutine** est un **processus léger** géré par le « Go runtime ».

```Go
go f(x, y, z)
```
<!-- .element: class="big-code" -->

démarre une nouvelle **goroutine** exécutant

```Go
f(x, y, z)
```
<!-- .element: class="big-code" -->

L'évaluation de **f**, **x**, **y** et **z** est effectuée dans la goroutine actuelle et l'exécution de **f** est effectuée dans la nouvelle goroutine.

Notes:
La fonction main (et donc les fonctions appelées depuis main) s'exécute sur sa goroutine, qu’on peut appeler la goroutine principale.

Le “Go runtime” est un package de Go qui est lié statiquement à l’exécutable lors de l’édition des liens après la compilation.

C’est un package un peu spécial qui contient entre autre le GC, la gestion des Goroutines, etc.

Voir https://www.quora.com/How-does-the-Go-runtime-work
