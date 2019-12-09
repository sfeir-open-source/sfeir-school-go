# Méthodes
Go ne dispose pas des classes. Cependant, vous pouvez définir des méthodes sur les types.

Une méthode est une fonction avec un argument de récepteur spécial.

Le récepteur apparaît dans sa propre liste d'argument entre le mot clé func et le nom de la méthode.

Dans cet exemple, la méthode Abs possède un récepteur de type Vertex nommé v.

    type Vertex struct {
        X, Y float64
    }

    func (v Vertex) Abs() float64 {
        return math.Sqrt(v.X*v.X + v.Y*v.Y)
    }
    
Vous pouvez déclarer une méthode sur les types non-struct, aussi.

Dans cet exemple, nous voyons un type numérique MyFloat avec une méthode Abs.

Vous ne pouvez que déclarer une méthode avec un récepteur dont le type est défini dans le même paquet que la méthode. 
Vous ne pouvez pas déclarer une méthode avec un récepteur dont le type est défini dans un autre paquet (ce-qui comprend les types intégrés comme int).

    type MyFloat float64

    func (f MyFloat) Abs() float64 {
        if f < 0 {
            return float64(-f)
        }
        return float64(f)
    }
    
## Récepteurs de Pointeur
Vous pouvez déclarer des méthodes avec des récepteurs de pointeur.

Cela signifie que le type de récepteur a la syntaxe littérale *T pour un certain type T. (Aussi, T ne peut pas être lui-même un pointeur comme *int.)

Par exemple, la méthode Scale est ici défini sur *Vertex.

Les méthodes avec récepteurs de pointeur peuvent modifier la valeur à laquelle le récepteur pointes (comme Scale le fait ici). 
Comme les méthodes ont souvent besoin de modifier leur récepteur, les récepteurs de pointeur sont plus communs que les récepteurs de valeur.

Essayez de retirer le * de la déclaration de la fonction Scale sur la ligne 16 et observez comment le comportement du programme change.

Avec un récepteur de valeur, la méthode Scale fonctionne sur une copie de la valeur original Vertex. 
(Ceci est le même comportement que pour toute autre argument de la fonction.) 
La méthode Scale doit disposer d'un récepteur de pointeur pour modifier la valeur Vertex déclarée dans la fonction main.

    type Vertex struct {
        X, Y float64
    }

    func (v Vertex) Abs() float64 {
        return math.Sqrt(v.X*v.X + v.Y*v.Y)
    }

    func (v *Vertex) Scale(f float64) {
        v.X = v.X * f
        v.Y = v.Y * f
    }