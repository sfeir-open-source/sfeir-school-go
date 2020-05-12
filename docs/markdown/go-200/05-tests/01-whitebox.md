<!-- .slide: class="with-code-bg-dark" -->

# Tests simples

## Le test peut être écrit dans le même package que ce que l’on teste ⇒ white box tests

```go
package example
import "testing"

func TestInternalStuff(t *testing.T) {
    // Test non-exported 'abs' function
    if abs(-3) != 3 {
        t.Error("Wrong result")
    }
}
```

Notes:
OFU

Test white box ⇒ nécessaire car sur un package entier, c’est compliqué de tester seulement les méthodes exportées

Pas de Mockito comme en java pour créer facilement des mocks
