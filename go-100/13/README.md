# Structures

Une struct est une collection de champs. Les champs de la structure sont accessibles à l'aide d'un point.

Une structure littérale dénote une valeur de structure nouvellement allouée en énumérant les valeurs de ses champs.

Vous pouvez lister seulement un sous-ensemble de champs en utilisant la syntaxe Nom:. (Et l'ordre des champs nommés n'est pas pertinent.)

    type Vertex struct {
        X, Y int
    }

    var (
        v1 = Vertex{1, 2}  // has type Vertex
        v2 = Vertex{X: 1}  // Y:0 is implicit
        v3 = Vertex{}      // X:0 and Y:0
    )