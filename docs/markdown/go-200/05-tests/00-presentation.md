<!-- .slide: class="with-code-bg-dark" -->

# Tests simples

- Package **testing** pour écrire les tests automatisés
- La commande **go test** pour exécuter les tests
- Un fichier de test doit se terminer par **\_test.go**

```go
package main
import "testing"
func TestDummy(t *testing.T) {
    if 4/2 != 2 {
        t.Error("4/2 should be equal to 2")
    }
}
```

```shell
$ go test

PASSok golang-200/dummy 0.012s
```

Notes:
OFU

Package “testing” fournit tout ce qu’il faut pour écrire des Tests Auto

La commande “go test” exécute tous les tests dans le répertoire courant et les sous répertoires

⇒ un test est identifié par un fichier se terminant par \_test.go et une fonction commençant par Test et prenant \*testing.T en paramètre.

Les fichier suffixés par \_test.go sont exclus de la compilation lors d’un build

Il existe des libs d’assertion
