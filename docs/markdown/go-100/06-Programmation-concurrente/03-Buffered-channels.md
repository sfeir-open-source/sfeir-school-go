<!-- .slide: class="with-code" -->

# Concurrence - 31

## Channels avec buffer

Les **channels** peuvent avoir un **buffer**.
Il faut indiquer la taille du buffer en second argument de make pour initialiser un channel avec buffer :

```Go
ch := make(chan int, 100)
```
<!-- .element: class="big-code" -->

L'envoi à un channel avec un buffer **bloque** uniquement **lorsque le buffer est plein**.
La réception **bloque lorsque le buffer est vide**.
