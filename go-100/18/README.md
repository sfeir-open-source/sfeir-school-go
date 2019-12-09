# Maps

Une map associe clés et valeurs.

Les maps doivent être créés avec make avant utilisation ; la map nil est vide et ne peut pas être assignée.

Les littéraux de maps sont comme des littéraux de struct, mais les clés sont nécessaires :

    var m = map[string]Vertex{
        "Bell Labs": Vertex{
            40.68433, -74.39967,
        },
        "Google": Vertex{
            37.42202, -122.08408,
        },
    }

## Manipulations de maps

Insérer ou mettre à jour un élément de map m :

    m[key] = elem
Récupérer un élément :

    elem = m[key]
Supprimer un élément :

    delete(m, key)
Test qu'une clé est présente avec une affectation de deux valeurs :

    elem, ok = m[key]
Si key est dans m , ok est true. Sinon, ok est false

Si key n'est pas dans la map, alors elem est la valeur zéro pour le type d'élément de la map.

Note : si elem ou ok ont pas encore été déclarée, vous pouvez utiliser une déclaration courte :

    elem, ok := m[key]