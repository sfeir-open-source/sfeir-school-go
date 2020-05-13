<!-- .slide: class="with-code" -->

# Benchmarks

**A ne jamais faire** dans un benchmark :

- changer les paramètres entre deux itérations

```go
func BenchmarkFibWrong(b *testing.B) {
    for n := 0; n < b.N; n++ {
        Fib(n)
    }
}
```

<!-- .element: class="big-code" -->

- utiliser le nombre de “runs” comme paramètre

```go
func BenchmarkFibWrong2(b *testing.B) {
    Fib(b.N)
}
```

<!-- .element: class="big-code" -->

Notes:
OFU

Ca ne convergera jamais ! Impossible d’établir une moyenne.
