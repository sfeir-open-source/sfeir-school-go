<!-- .slide: class="with-code" -->

# Concurrence

## On va créer notre première structure (objet)

```go
// Statistics is the worker to persist the request statistics
type Statistics struct {
    statistics chan bool
    counter uint32
    start time.Time
    loggingPeriod time.Duration
}
```

<!-- .element: class="big-code" -->

Notes:
SFR

structure de base

type de données uint8 ⇒ complexe128
