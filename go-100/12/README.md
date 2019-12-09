# Les types

Il est possible de déclarer un nouveau type à l'aide du mot clé type et d'un type sous-jacent

    type nom typeSousJacent
    
Exemples :

    type distance float64
    type HttpStatusCode int16

Il est possible d’effectuer les même opérations que sur le type sous-jacent, mais seulement entre variables du même type.
C’est très utile pour éviter les erreurs d’homogénéité dans une formule.
Il est toujours possible d'utiliser la conversion de type pour effectuer des opérations :

    type Distance int64
    type Duration int64
    var d Distance = 300
    var t Duration = 60
    x := d * t // erreur de compilation
    x := float64(d) * float64(t)  // ok
