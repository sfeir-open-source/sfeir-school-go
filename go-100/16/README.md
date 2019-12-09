# Slices
Un Slice pointe vers un tableau de valeurs et possède une longueur.

[]T est un slice avec des éléments de type T.

len(s) retourne la longueur du slice s.

Les slices peuvent contenir tout type, y compris d'autres slices :

    game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
    
## Slices redécoupés
Les slices peuvent être redécoupés, créant ainsi un nouveau slice qui pointe vers le même tableau.

L'expression

    s[lo:hi]
récupère les éléments du slice de lo à hi-1 inclus. Ainsi

    s[lo:lo]
est vide et

    s[lo:lo+1]
possède un élément.

## Créer des Slices
Les Slices peuvent être créés avec la fonction make. Cela fonctionne en allouant un tableau vide et en retournant un slice qui référence ce tableau :

    a := make([]int, 5)  // len(a)=5
Pour spécifier une capacité, passer un troisième argument à make :

    b := make([]int, 0, 5) // len(b)=0, cap(b)=5

    b = b[:cap(b)] // len(b)=5, cap(b)=5
    b = b[1:]      // len(b)=4, cap(b)=4
    
## Slices nuls
La valeur nulle d'un slice est nil.

Un slice nil possède une longueur et une capacité de zéro.

## Ajout d'éléments à un slice
Il est courant d'ajouter de nouveaux éléments à un slice, et ainsi Go offre une fonction intégrée append. La documentation décrit append comme ceci :

    func append(s []T, vs ...T) []T

Le premier paramètre s de append est un slice de type T, et le reste sont des valeurs T à ajouter au slice.

La valeur résultante de append est un slice contenant tous les éléments du slice original et les valeurs fournies.

Si le tableau de support de s est trop petit pour contenir toutes les valeurs données, un plus gros tableau sera alloué. Le slice retourné pointera vers le tableau nouvellement alloué.