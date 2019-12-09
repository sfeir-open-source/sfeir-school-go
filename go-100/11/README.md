# Defer

Une déclaration "defer" reporte l'exécution d'une fonction jusqu'à ce que la fonction environnante retourne.

Les arguments de l'appel différé sont évalués immédiatement, mais l'appel de fonction n'est pas exécuté jusqu'à ce que la fonction environnante retourne.

Les appels de fonction différés sont poussés sur une pile. Quand une fonction retourne, ses appels différés sont exécutées dans l'ordre dernier entré, premier sorti.

Les appels de fonction différés sont effectués même en cas de panique (panic).
"defer" peut donc être utilisé pour rattraper des erreurs inattendues et les traiter, dans le même esprit qu'un try/catch en java ou C#, sauf
que ces derniers ont tendance à complexifier la lecture du code contrairement à "defer".