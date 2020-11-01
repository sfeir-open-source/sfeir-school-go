<!-- .slide: class="with-code" -->

# Concurrence - 32

## Fermeture de channels

Un expéditeur peut fermer un canal pour indiquer que plus aucune valeur ne sera envoyée :

```Go
close(ch)
```
<!-- .element: class="big-code" -->

Les récepteurs peuvent vérifier si un canal a été fermé :

```Go
v, ok := <-ch
```
<!-- .element: class="big-code" -->

**ok** est false si le canal est fermé et qu’il n'y a plus de valeur à recevoir.

##==##

# Concurrence - 32

## Fermeture de channels

Attention :

- **Seul l'expéditeur doit fermer un canal**, jamais le récepteur.
  L'envoi sur un canal fermé provoque une panique.
- Les canaux ne sont pas comme les fichiers, vous n'avez généralement **pas besoin de les fermer**.
  La fermeture n'est nécessaire que lorsque le récepteur doit être informé qu'il n'y a plus de valeurs à venir, comme mettre fin à la boucle d'un range.
