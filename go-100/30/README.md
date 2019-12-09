# Goroutines
Une goroutine est un processus léger géré par le « Go runtime ».

    go f(x, y, z)
    
démarre une nouvelle goroutine exécutant

    f(x, y, z)

L'évaluation de f, x, y et z arrive dans la goroutine actuelle et l'exécution de f arrive dans la nouvelle goroutine.

Les goroutines s'exécutent dans le même espace d'adressage, de sorte que les accès à la mémoire partagée doivent être synchronisés. 
Le paquet sync fournit des primitives utiles, bien que vous n'aurez pas beaucoup besoin d'y Accéder car Go a d'autres primitives plus pratiques.