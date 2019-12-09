# Canaux (Channels)
Les canaux sont des conduits typés à travers lesquels vous pouvez envoyer et recevoir des valeurs avec l'opérateur de canal, <-.

    ch <- v    // Envoyer v sur le canal ch.
    v := <-ch  // Recevoir de ch, et
               // attribuer une valeur à v.
            
(Le flux de données vas dans le sens de la flèche.)

Comme les cartes et les tranches, les canaux doivent être créés avant l'utilisation :

    ch := make(chan int)
    
Par défaut, l'envoi et la réception bloque jusqu'à ce que l'autre côté soit prêt. 
Cela permet de synchroniser les goroutines sans verrous explicites ou variables de condition.