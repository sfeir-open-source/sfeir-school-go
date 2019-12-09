# Select

La déclaration select permet à une goroutine d'attendre sur plusieurs opérations de communication simultanément.

Un select bloque jusqu'à ce que l'un de ses cas puisse s'exécuter, puis il l'exécute. Il en choisit un au hasard si plusieurs sont prêts.