# Les fonctions
Une fonction est déclarée à l'aide du mot clé func.

    func nomDeLaFonction() {
    }

Une fonction peut prendre zéro ou plusieurs arguments.

    func nomDeLaFonction(param1 int, param2 string) {
    }

De même, une fonction peut retourner une ou plusieurs valeurs.

    func nomDeLaFonction(param1 int, param2 string) int {
    }
    func nomDeLaFonction(param1 int, param2 string) (int, error) {
    }

Les valeurs de retour peuvent aussi être nommées.

    func maFonction3(param1 int, param2 string) (result int, err error) {
    }

Notez que le type vient après l'identifiant de la variable.