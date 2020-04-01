<!-- .slide: class="with-code" -->

# Concurrence - 34

## For range sur les channels

La boucle for `i := range c` reçoit les valeurs du canal jusqu'à ce qu'il soit fermé.

```Go
for i := range c { fmt.Println(i)}
```
