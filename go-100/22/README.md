# Méthodes et pointeur d'indirection
En comparant, vous remarquerez peut-être que des fonctions avec un argument pointeur doivent prendre un pointeur :

    var v Vertex
    ScaleFunc(v)  // Compile error!
    ScaleFunc(&v) // OK

tandis que les méthodes avec des récepteurs de pointeur prennent une valeur ou un pointeur comme récepteur quand ils sont appelés :

    var v Vertex
    v.Scale(5)  // OK
    p := &v
    p.Scale(10) // OK

Pour la déclaration v.Scale(5), même si v est une valeur et non un pointeur, la méthode avec le récepteur de pointeur est appelée automatiquement.
Autrement dit, à titre de commodité, Go interprète la déclaration v.Scale(5) comme (&v).scale(5) car la méthode Scale possède un récepteur de pointeur.