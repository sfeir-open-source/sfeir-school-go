# if

En Go, la déclaration if s'écrit :

    if condition {
    }

La condition n'a pas besoin d'être entouré par des 
parenthèses ( ) mais les accolades { } sont obligatoires.

L'instruction if peut commencer par une brève déclaration à exécuter avant la condition.

Les variables déclarées au niveau d'un if ont seulement une portée jusqu'à la fin de l'instruction if, mais elles
sont également disponibles à l'intérieur des blocs else :

    if v := expression; condition {
        fmt.Println(v)
    } else {
        fmt.Println(v)
    }