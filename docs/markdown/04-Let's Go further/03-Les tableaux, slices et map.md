<!-- .slide:-->

# Pour aller plus loin - 15


Le type [n]T est un tableau de n valeurs de type T.

L’expression :
var a [10]int
déclare une variable a comme un tableau de dix entiers.


Les tableaux


##==##
<!-- .slide:-->

# Pour aller plus loin - 15


La longueur d'un tableau fait partie de son type.

⇒ les tableaux ne peuvent pas être redimensionnés.


Les tableaux


Notes:
La longueur d'un tableau fait partie de son type, de sorte que les tableaux ne peuvent pas être redimensionnés. Cela semble limitatif, mais ne vous inquiétez pas ; Go offre un moyen pratique de travailler avec des tableaux.




##==##
<!-- .slide:-->

# Pour aller plus loin - 16


Un slice pointe vers un tableau de valeurs et possède une longueur.
[]T est un slice avec des éléments de type T.

len(s) retourne la longueur du slice s.
cap(s) retourne la capacité du slice s.
La valeur zéro d’un slice est nil.



Les slices


##==##
<!-- .slide:-->

# Pour aller plus loin - 16


Les slices peuvent être créés avec la fonction make. Cela alloue un tableau vide et retourne un slice qui référence ce tableau :

a := make([]int, 5)  // len(a)=5




Les slices


##==##
<!-- .slide:-->

# Pour aller plus loin - 16


Les slices peuvent être redécoupés, créant ainsi un nouveau slice qui pointe vers le même tableau.

L’expression :
s[low:high]
récupère les éléments du slice de low à high-1 inclus (ou high exclus).



Les slices


##==##
<!-- .slide:-->

# Pour aller plus loin - 16


Les slices


var a = make([]byte, 5)len(a) == 5cap(a) == 5


var b = a[2:4]len(b) == 2cap(b) == 3


![](./images/slice-2-array.png)

![](./images/slice-2-array-decoupe.png)

##==##
<!-- .slide:-->

# Pour aller plus loin - 16


Pour spécifier une capacité, passer un troisième argument à make :

b := make([]int, 0, 5) // len(b)=0, cap(b)=5b = b[:cap(b)]             // len(b)=5, cap(b)=5b = b[1:]                      // len(b)=4, cap(b)=4



Les slices


##==##
<!-- .slide:-->

# Pour aller plus loin - 16


Il est courant d'ajouter de nouveaux éléments à un slice, et ainsi Go offre une fonction intégrée append.

func append(s []T, vs ...T) []T
Le premier paramètre s de append est un slice de type T, et les suivants sont des valeurs T à ajouter au slice.



Les slices


Notes:
La valeur résultante de
append
 est une tranche contenant tous les éléments de la tranche original et les valeurs fournis.

Si le tableau de support de
s
 est trop petit pour contenir toutes les valeurs données un plus gros tableau sera alloué. La tranche retourné pointera vers le tableau nouvellement allouée.





##==##
<!-- .slide:-->

# Pour aller plus loin - 17


La forme range de la boucle for permet d’itérer sur tous les éléments d’un slice, d’un tableau, ou d’une map.

Deux valeurs sont renvoyées pour chaque itération :
l'indice
une copie de l'élément à cet indice.


For … range


##==##
<!-- .slide:-->

# Pour aller plus loin - 17


for i, v := range s {}

Est équivalent à :
for i := 0; i < len(s); i++ {	v := s[i] // s[i] est ici modifiable}


For … range


##==##
<!-- .slide:-->

# Pour aller plus loin - 17


On peut omettre l’indice en le remplaçant par “_” :
for _, v := range s { }

On peut aussi omettre la valeur.
for i := range s { }



For … range


##==##
<!-- .slide:-->

# Pour aller plus loin - 17 - exercice


Écrire les fonctions indexOf et split.


Les slices et le range


Notes:
-> 2h



##==##
<!-- .slide:-->

# Pour aller plus loin - 18


Une map est une liste de clés-valeurs. Une map se déclare :
map[typeClé]typeValeur

Les maps doivent être créées avec make. Leur valeur-zéro est nil.
m = make(map[string]Vector)



Maps


##==##
<!-- .slide:-->

# Pour aller plus loin - 18


Comme les structs, une map peut être instanciée littéralement. Il faut cependant indiquer la valeur de chaque clé :
var m = map[string]Vector{	"Bell Labs": Vector{40.68433, -74.39967},	"Google": Vector{37.42202, -122.08408},}




Maps


##==##
<!-- .slide:-->

# Pour aller plus loin - 18


Affectation : m[key] = elem
Suppression : delete(m, key)
Récupération :
elem = m[key] //elem prend la valeur-zéro de son type si la clé n'est pas présente.
elem, ok = m[key] //ok est true si la clé est présente, sinon false.


Maps


Notes:
Insérer ou mettre à jour un élément de carte
m
 :

m[key] = elem

Récupérer un élément :

elem = m[key]

Supprimer un élément :

delete(m, key)

Test qu'une clé est présente avec une affectation de deux valeurs :

elem, ok = m[key]

Si
key
 est dans
m , ok
 est
true
. Sinon,
ok
 est
false

Si
key
 n'est pas dans la carte, alors
elem
 est la valeur zéro pour le type d'élément de la carte.

Note :
 si
elem
 ou
ok
 ont pas encore été déclarée, vous pouvez utiliser une déclaration courte :

elem, ok := m[key]





##==##
<!-- .slide:-->

# Pour aller plus loin - 19


En Go, les fonctions sont aussi des valeurs. Elles peuvent être passés comme toute autre valeur.
Les valeurs de fonction peuvent être utilisées comme arguments de fonctions ou valeurs de retour. On dit que les fonctions sont first class citizen en Go



Un peu plus sur les fonctions


##==##
<!-- .slide:-->

# Pour aller plus loin - 19


Les fonctions de Go peuvent être des closures. Une closure (fermeture) est une valeur de la fonction qui fait référence à des variables à partir de l'extérieur de son corps.
La fonction peut accéder et assigner les variables référencées, dans ce sens, la fonction est «liée» aux variables.


Closures


##==##
<!-- .slide:-->

# Pour aller plus loin - 19 - exercice


Implémentez la fonction multiplicateurPar pour qu'elle retourne une fonction qui multiplie son argument par un nombre donné.

//fonction qui retourne une fonction qui retourne un int.
func multiplicateurPar(x int) func(int) int {}




Closures


##==##
<!-- .slide:-->

# Pour aller plus loin - 20


Il est possible en Go de déclarer plusieurs variables sur une même ligne.

a, b, c := 1, "toto", 8




Déclarations multiples


##==##
<!-- .slide:-->

# Pour aller plus loin - 20


De même, il est possible d’assigner plusieurs variables sur une même ligne.
C’est très pratique pour effectuer un swap :

x, y = y, x



Assignations multiples

