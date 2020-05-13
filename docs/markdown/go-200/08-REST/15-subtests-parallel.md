<!-- .slide: class="with-code" -->

# Subtests

## Les subtests peuvent être exécutés en parallèle

```go
func TestFoo(t *testing.T) {
    t.Run("A=1", func(t *testing.T) {
        t.Parallel()
    /// ...
    })

    t.Run("A=2", func(t *testing.T) {
        t.Parallel()
    // ...
    })
}
```

<!-- .element: class="big-code" -->

Notes:
OFU

Pour le coup, les tests doivent
impérativement
être indépendants !
