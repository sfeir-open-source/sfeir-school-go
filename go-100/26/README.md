# Attention à ne pas confondre T et *T

En général, toutes les méthodes sur un type donné T doivent avoir soit un récepteur de valeur (T) soit un récepteur de pointeur (*T), mais pas un mélange des deux.

Sinon, T implémenterait des interfaces tandis que *T en implémenterait d’autres, ce qui deviendrait pénible à utiliser.
