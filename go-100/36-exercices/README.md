# Exercice: Requêtes HTTP avec goroutine et select

Modifiez le programme pour que les erreurs soient envoyées sur un canal spécifique (chan error) et récupérées via un select dans la fonction main.