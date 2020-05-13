<!-- .slide: class="with-code" -->

# Subtests

## Les subtests pour mieux organiser et regrouper les tests

```go
func TestFoo(t *testing.T) {
    // ...
    t.Run("A=1", func(t *testing.T) { ... })
    t.Run("A=2", func(t *testing.T) { ... })
    t.Run("B=1", func(t *testing.T) { ... })
    // ...
}
```

<!-- .element: class="big-code" -->

```shell
$ go test -run Foo    # Run top-level tests matching "Foo".
$ go test -run Foo/A= # Run subtests of Foo matching "A=".
$ go test -run /A=1   # Run all subtests of a top-level test matching "A=1".
```

<!-- .element: class="big-code" -->

Notes:
OFU

permet d’effectuer des tests dans un ordre bien précis (setup/teardown)

Mais attention ⇒ c’est quand même mieux de faire des tests indépendants :

pour pouvoir les exécuter séparément

pour pouvoir les exécuter en parallèle
