<!-- .slide: class="with-code" -->

# Parsin des arguments

## Action principale et lancement

```go
package main
import "github.com/urfave/cli/v2"
var port = 8020

func main() {
    // […]
    // main action
    // sub action are possible also
    app.Action = func(c *cli.Context) error {
        // parse parameters
        fmt.Printf("port : %d\n", port)
        return nil
    }

    // […]
    // run the app
    err := app.Run(os.Args)
    if err != nil {
        fmt.Errorf("Run error %q\n", err)
    }
}
```
