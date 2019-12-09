# L'interface vide
Le type d'interface qui ne spécifie aucune méthode est connu comme l'interface vide :

    interface{}
    
Une interface vide peut contenir des valeurs de tout type. (Chaque type implémente au moins zéro méthode.)

Les interfaces vides sont utilisés par du code qui gère les valeurs de type inconnu. Par exemple, fmt.Print prend un nombre d'arguments de type interface{}.