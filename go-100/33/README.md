# Fermeture d'un canal

Un expéditeur peut close un canal pour indiquer qu'il n'y a plus de valeurs qui seront envoyés. 
Les récepteurs peuvent vérifier si un canal a été fermée par l'attribution d'un deuxième paramètre à l'expression de réception :

    v, ok := <-ch
    
ok est false si il n'y a plus de valeurs à recevoir et le canal est fermé.