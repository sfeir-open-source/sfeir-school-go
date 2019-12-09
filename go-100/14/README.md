# Pointeurs
Go comporte des pointeurs explicites. Un pointeur contient l'adresse mémoire d'une variable.

Le type *T est un pointeur vers une valeur de T. Sa valeur zéro est nil.

    var p *int

L'opérateur & génère un pointeur vers son opérande.

    i := 42
    p = &i

L'opérateur * dénote la valeur sous-jacente du pointeur.

    fmt.Println(*p) // lire i par le pointeur p
    *p = 21         // définir i par le pointeur p

Ceci est connu en tant que « déréférencement » ou « indirection ».

Contrairement à C, Go n'a aucune arithmétique de pointeur. L'arithmétique de pointeur est très (trop ?) puissante, mais elle complexifie le code.
Go portant une grande importance à la simplicité, il est logique qu'il ne permette pas d'effectuer des opérations sur les pointeurs.