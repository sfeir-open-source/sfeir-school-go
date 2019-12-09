# Range
La forme range de la boucle for effectue une itération sur un slice ou une map.

Lors de l'itération sur un slice, deux valeurs sont renvoyées pour chaque itération. La première est l'indice, et la seconde est une copie de l'élément à cet indice.

    for i, value := range s {
	}

Vous pouvez sauter l'indice ou la valeur en l'attribuant à _.

Si vous ne souhaitez que l'indice, retirez ", value" entièrement.

    //Indice uniquement
    for i := range s {
	}
    //Valeur uniquement
	for _, value := range s {
	}