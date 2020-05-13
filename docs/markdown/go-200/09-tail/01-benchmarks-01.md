<!-- .slide: class="with-code-bg-dark" -->

# Benchmarks

- Le package **testing** fournit tout ce qu’il faut pour créer des **benchmarks**
- Une fonction de benchmark :
  - commence par **Benchmark**
  - prend en paramètre le nombre d'exécutions à effectuer **b.N**
  - peut être exécutée plusieurs fois avec des valeurs différentes de **b.N**

```go
func BenchmarkFib10(b *testing.B) {
    // run the Fib function b.N times
    for n := 0; n < b.N; n++ {
        Fib(10)
    }
}
```

```shell
$ go test -bench=.

PASS
BenchmarkFib10 5000000 509 ns/op
ok github.com/Sfeir/fib 3.084s
```

![](./assets/go-200/images/speedometer.gif)

<!-- .element style="position:absolute; right: 50px; top: 50px; max-width: 380px" -->

Notes:
OFU

peut être exécuté plusieurs fois jusqu’à avoir un résultat suffisamment stable (moyenne précise)
