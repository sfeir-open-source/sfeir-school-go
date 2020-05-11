<!-- .slide: class="with-code-bg-dark" -->

# Parsing des arguments

## std lib flag

Go possède sa propre lib de parsing de la ligne de commande :

https://golang.org/pkg/flag

<br>

```go
package main
import ( "fmt" "flag")

func main() {
    var flagvar int
    flag.IntVar(&flagvar, "intp", 42, "help message for int flag")
    flag.Parse()
    fmt.Println("flagvar has value ", flagvar)
}
```

```shell
$ go run main.go -intp 314
flagvar has value 314
```
