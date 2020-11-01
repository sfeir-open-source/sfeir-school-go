<!-- .slide: class="with-code" -->

# Tests simples

## Pour tester notre agrégateur de statistique

```go
func TestStatistics(t *testing.T) {
    statistics := NewStatistics(2 _ statPeriod)

    // other go routine incrementing the counter
    go func() {
        statistics.PlusOne()
    }()

    time.Sleep(statPeriod)

    if statistics.counter != 1 {
        t.Errorf("Wrong count %d ; expected %d", statistics.counter, 1)
    }

    time.Sleep(2 * statPeriod)

    if statistics.counter != 0 {
        t.Errorf("Wrong count %d ; expected %d", statistics.counter, 0)
    }
}
```
