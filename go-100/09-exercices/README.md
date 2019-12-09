# Exercice : Syracuse

Implémentez l'algorithme de la conjecture de Syracuse :

    soit n un nombre entier > 1,
    tant que n != 1, faire
        si n est pair
            n = n/2
        sinon
            n = 3n + 1
        fin si
    fin
    
On ajoutera en paramètre de la fonction un nombre maximum d'itérations à effectuer. 
La fonction devra retourner un booléen indiquant si le nombre 1 a été atteint (avant que le nombre max d'itérations ne soit atteint).
Elle retournera aussi le nombre d'itérations effectuées.

Note : pour savoir si un nombre est pair, on peut utiliser le modulo % (ex: n%2 == 0).