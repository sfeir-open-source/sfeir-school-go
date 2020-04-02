<!-- .slide: class="sfeir-bg-white-3" -->

# Les bases - 11

## Defer

Une déclaration defer reporte l'exécution d'une fonction jusqu'à ce que la fonction environnante retourne.

Les arguments de l'appel différé sont évalués immédiatement, mais l'appel de fonction n'est pas exécuté jusqu'à ce que la fonction environnante retourne.


##==##
<!-- .slide: class="with-code" -->

# Les bases - 11

## Defer

Les appels de fonction différés sont poussés sur une pile (FILO) et sont effectués même en cas de panique.

```Go
func main() {
	defer fmt.Print("world\n")
	defer fmt.Print(", ")
	fmt.Print("Hello")
}
```
<!-- .element: class="big-code" -->


Notes:
Les appels de fonction différés sont poussés sur une pile. Quand une fonction retourne, ses appels différés sont exécutées dans l'ordre dernier entré, premier sorti.



