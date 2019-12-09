# Select default

Le cas default dans un select est exécuté si aucun autre cas n'est prêt et permet donc d’éviter le blocage du select.

    select {
    case i := <-c:
        // utiliser i
    default:
        // exécuté si rien ne peut être lu de 'c'
    }