<!-- .slide: class="with-code-bg-dark" -->

# Parsing des arguments

## lib urfave cli

Mais on est barbu·e ou on l’est pas…

!["Dave Cheney" center hm-200](./assets/go-200/images/DaveCheney.webp)

- https://github.com/urfave/cli/v2 (ex-codegangsta)
- v2 finalisée

```go
package main
import "github.com/urfave/cli/v2"
var port = 8020
func main() {
    // new app
    app := cli.NewApp()
    app.Name = "mytodolist"
    app.Usage = "mytodolist service launcher"
    // […]
}
```

Notes:
définition d’une nouvelle application

variables positionnées à la compilation

plusieurs main possible

dave cheney

il existe aussi Viper
