# Switch

Le switch s'écrit comme cela :

    switch expression {
    case a:
        ...
	case b:
        ...
	default:
        ...
    }

Un cas ("case") est terminé automatiquement, sauf s'il finit par une déclaration fallthrough.

Le switch évalue les cas de haut en bas, en s'arrêtant quand un cas correspond.

(Par exemple,

    switch i {
    case 0:
    case f():
    }
n'appelle pas f si i==0.)