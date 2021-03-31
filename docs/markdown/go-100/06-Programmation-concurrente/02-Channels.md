<!-- .slide: class="with-code" -->

# Concurrence - 30

## Channels (canaux)

Les **channels** sont des conduits typés à travers lesquels vous pouvez envoyer et recevoir des valeurs avec l'opérateur de canal, **<-**

```Go
ch <- v // Envoyer v sur le canal ch.
v := <- ch // Recevoir de ch, et attribuer la valeur à v.
```

<!-- .element: class="big-code" -->

(Le flux de données va dans le sens de la flèche.)

##==##

<!-- .slide: class="with-code" -->

# Concurrence - 30

## Channels (canaux)

Les **channels** doivent être créés avant l'utilisation :

```Go
    ch := make(chan int)
```

<!-- .element: class="big-code" -->

Par défaut, l'envoi et la réception bloquent jusqu'à ce que l'autre côté soit prêt. Cela permet de synchroniser les goroutines sans verrous explicites ou variables de condition.
