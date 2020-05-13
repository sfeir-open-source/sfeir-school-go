<!-- .slide: class="with-code" -->

# Concurrence

## Une méthode pour la go routine principale d'agrégation

```go
func (sw *Statistics) run() {
    ticker := time.NewTicker(sw.loggingPeriod)
    for {
        select {
            case stat := <-sw.statistics:
                logger.WithField("stat", stat).Debug("new count received")
                sw.counter++
            case now := <-ticker.C:
                elapsed := now.Sub(sw.start)
                logger.WithField("since", elapsed).WithField("count", sw.counter).Warn("monitoring")
                // reset
                sw.counter = 0
                sw.start = now
        }
    }
}
```

Notes:
SFR

méthode

minuscule car privée

select pour multiplexage double channel
