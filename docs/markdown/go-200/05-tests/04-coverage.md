<!-- .slide: class="with-code-bg-dark" -->

# Tests simples

## Pour générer le rapport de coverage du code

```shell
$ go test -cover -coverprofile=testcover.out
$ go tool cover -html=testcover.out
```
