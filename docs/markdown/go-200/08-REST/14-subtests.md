<!-- .slide: class="with-code-bg-dark" -->

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

```shell
$ go test -run Foo    # Run top-level tests matching "Foo".
$ go test -run Foo/A= # Run subtests of Foo matching "A=".
$ go test -run /A=1   # Run all subtests of a top-level test matching "A=1".
```

Notes:
OFU

permet d’effectuer des tests dans un ordre bien précis (setup/teardown)

Mais attention ⇒ c’est quand même mieux de faire des tests indépendants :

pour pouvoir les exécuter séparément

pour pouvoir les exécuter en parallèle
