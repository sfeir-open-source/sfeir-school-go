# Assertions de type
Un assertions de type donne accès à la valeur sous-jacente concrète d'une valeur d'interface.

    t := i.(T)
    
Cette instruction affirme que la valeur d'interface i contient le type concret T et assigne la valeur sous-jacent T à la variable t.

Si i ne détient pas T, la déclaration va déclencher une panique.

Pour tester si une valeur d'interface détient un type spécifique, une assertion de type peut retourner deux valeurs: la valeur sous-jacente et une valeur booléenne qui indique si l'affirmation réussi.

    t, ok := i.(T)

Si i détient T, alors t sera la valeur sous-jacente et ok sera vrai.

Si non, ok sera faux et t sera la valeur zéro de type T, et aucune panique se produit.

Notez la similitude entre cette syntaxe et celle de la lecture d'une carte.