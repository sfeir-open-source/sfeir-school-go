# Concurrence - 34

## Select

La déclaration **select** permet à une goroutine d'attendre sur plusieurs opérations de communication simultanément.

```Go
select {
    case c <- x: // attend de pouvoir envoyer x sur 'c'
        x, y = y, x+y
    case <- quit: // attend de recevoir depuis 'quit'
        fmt.Println("quit")
}
```
<!-- .element: class="big-code" -->

##==##

<!-- .slide: class="with-code" -->

# Concurrence - 34

## Select

Un **select** bloque jusqu'à ce que l'un de ses cas puisse s'exécuter, puis il l'exécute. Il en choisit un au hasard si plusieurs sont prêts.

##==##

<!-- .slide: class="with-code" -->

# Concurrence - 35

## Select

Le cas **default** dans un **select** est exécuté si aucun autre cas n'est prêt et permet donc d’éviter le blocage du select.

```Go
select {
    case i := <-c:
 // utiliser i
    default:
 // exécuté si rien ne peut être lu de 'c'
}
```
<!-- .element: class="big-code" -->
