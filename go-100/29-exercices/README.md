# Exercice : les interfaces

Créer un type FmtLogger qui implémente Logger

Puis utiliser le logger pour écrire ce que contient la panique (r) dans la fonction recoverPanic.
Vous devrez gérer le cas où r est un string, où r est une error, et un cas par défaut où on loggera "PANIC: unknown error".

Finalement, vous devrez appeler la fonction recoverPanic de façon à ce qu'elle soit appelée même en cas de panique.
Si vous ne savez pas comment faire, jetez un coup d'oeil à la section n°11.