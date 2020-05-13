<!-- .slide: class="with-code" -->

# REST

## Le middleware de statistique

```go
func NewStatisticsMiddleware(duration time.Duration) *StatisticsMiddleware {
    return &StatisticsMiddleware{
        Stat: statistics.NewStatistics(duration),
    }
}

func (sm *StatisticsMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    // [...]
    sm.Stat.PlusOne()
    next(rw, r)
    // [...]
}
```
