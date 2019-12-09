# Identifiant exporté

## Comment exporter une variable, une fonction, un type, etc.

En Go, un identifiant (qu'il fasse référence à une variable, un type, une fonction, une propriété, etc.) est exporté s'il commence par une majuscule.

Lors de l'importation d'un package, vous pouvez seulement faire référence à ses identifiants exportés. Tous les identifiants «non exportées» sont inaccessibles depuis l'extérieur du package.

Pi est un identifiant exporté, comme l'est PI. L'identifiant pi n'est pas exporté.

## Essayez

Compilez le code. Remarquez le message d'erreur.

Pour corriger l'erreur, renommez math.pi en math.Pi et essayez à nouveau.

Note: dans Visual Studio Code, vous pouvez voir le message d'erreur directement depuis l'éditeur, sans avoir à le compiler.