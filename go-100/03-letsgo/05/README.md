# Les variables
Une variable est déclarée à l'aide du mot clé `var` suivi du nom de la variable puis de son type ou de sa valeur initiale.

## Exemples de déclarations
    var nom type = expression
    var nom = expression
    var nom type

## Déclaration courte
A l'intérieur d'une fonction, la notation `:=` peut être utilisée au lieu de `var`.

    nom := expression

est strictement identique à

    var nom = expression

## Déclarations et assignations multiples

En Go, il est possible de déclarer ou d'assigner plusieurs variables sur une seule ligne.

C'est très pratique pour effectuer un swap :

    x, y = y, x
