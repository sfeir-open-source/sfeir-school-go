# Les valeurs de fonction
Les fonctions sont aussi des valeurs. Elles peuvent être passés comme toute autre valeur.

Les valeurs de fonction peuvent être utilisées comme arguments de fonctions et valeurs de retour.

## Fermetures de fonction (Closures)
Les fonctions de Go peuvent être des fermetures. Une fermeture est une valeur de la fonction qui fait référence à des variables à partir de l'extérieur de son corps.
La fonction peut accéder et assigner les variables référencées, dans ce sens, la fonction est «lié» aux variables.

Par exemple, la fonction de adder renvoie une fermeture. Chaque fermeture est liée à sa propre variable sum.

    func adder() func(int) int {
        sum := 0
        return func(x int) int {
            sum += x
            return sum
        }
    }

    func main() {
        pos, neg := adder(), adder()
        for i := 0; i < 10; i++ {
            fmt.Println(
                pos(i),
                neg(-2*i),
            )
        }
    }