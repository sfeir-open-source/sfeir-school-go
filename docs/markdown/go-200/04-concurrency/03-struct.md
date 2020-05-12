<!-- .slide: class="with-code-bg-dark" -->

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

Notes:
SFR

structure de base

type de données uint8 ⇒ complexe128
