<!-- .slide: class="with-code" -->

# Tests simples

## Ou bien dans le même package suffixé par \_test ⇒ black box tests

```go
package example_test
import (
    "testing"
    "example"
)

func TestExternalStuff(t *testing.T) {
    // Test exported function
    if example.Abs(-3) != 3 {
        t.Error("Wrong result")
    }
}
```

<!-- .element: class="big-code" -->
