# Les channels avec buffer
Les channels peuvent être bufferisées. Il suffit d'indiquer la longueur du buffer en second argument de make pour initialiser un canal avec buffer :

    ch := make(chan int, 100)
    
L'envoi à un canal avec buffer bloque uniquement lorsque le buffer est plein. La réception bloque lorsque le buffer est vide.

Modifiez l'exemple pour trop remplir le buffer et voir ce qui se passe.