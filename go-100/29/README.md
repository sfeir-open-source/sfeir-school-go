# Switch de type
Un switch de type est une construction qui permet d'effectuer plusieurs assertions de type en série.

Un switch de type est comme une instruction switch régulière, mais les cas d'un switch de type spécifient les types (pas les valeurs), et ces valeurs sont comparées au type concret de la valeur d'interface donnée.

    switch v := i.(type) {
    case T:
        // here v has type T
    case S:
        // here v has type S
    default:
        // no match; here v has the same type as i
    }
    
La déclaration dans un switch de type a la même syntaxe qu'une affirmation de type i.(T), mais le type spécifique de T est remplacé par le mot-clé type.

Cette instruction switch teste si la valeur d'interface i détient une valeur de type T ou S. Dans chacun des cas de T et S, la variable v sera de type T ou S respectivement, et détient la valeur détenue par i. Dans le cas par défaut (où il n'y a pas de correspondance), la variable v est de même type d'interface et de valeur que i.