# Récepteur nil

Si une méthode est appelée via un pointeur ayant la valeur nil, la méthode sera appelée avec un récepteur nil.

Dans certains langages, cela déclencherait une exception de pointeur nul, 
mais en Go il est courant d'écrire des méthodes qui gèrent avec élégance le fait d'être appelé avec un récepteur nil.