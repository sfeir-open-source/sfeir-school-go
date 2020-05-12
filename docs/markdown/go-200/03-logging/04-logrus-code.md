<!-- .slide: class="with-code-bg-dark" -->

# Logging

```go
package main
import ( "github.com/sirupsen/logrus" "os")

func main() {
    logrus.SetFormatter(&logrus.TextFormatter{
        ForceColors: true,
        FullTimestamp: true,
    })
    logrus.SetOutput(os.Stdout)
    logrus.SetLevel(logrus.DebugLevel)
    logrus
        .WithField("error", err)
        .Warn("error setting log level, using debug as default")
}
```

Exécution :

```shell
WARN[2017-03-04T23:37:34+01:00] error setting log level, using debug as default error="log an error"
```

Notes:
OFU

possibilité de créer un nouveau logger, mais logrus a un logger global par défaut
