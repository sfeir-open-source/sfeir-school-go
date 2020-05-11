<!-- .slide: class="with-code-bg-dark" -->

# Parsing des arguments

## Déclaration des flags

```go
package main
import "github.com/urfave/cli/v2"
var port = 8020

func main() {
    // […]
    // command line flags
    app.Flags = []cli.Flag{
        &cli.IntFlag{
            Value: port,
            Name: "port",
            Aliases: []string{"p"},
            Usage: "Set the listening port of the webserver",
            Destination: &port,
            EnvVars: []string{"APP_PORT"},
        },
    }
}
```
