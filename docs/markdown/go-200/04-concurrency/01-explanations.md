# Concurrence

- Problématique : compter le nombre de **requêtes par seconde** sur notre API
- Environnement : **multi-threadé** (Go routines)
- Contrainte : pas de **Mutex** pour ne pas plomber les **perfs**
- Solution apportée par Go : les canaux ! (**channels**)

<blockquote cite="https://www.youtube.com/watch?v=PAAkCSZUG1c&t=2m48s">
    <p>Don't communicate by sharing memory, share memory by communicating.</p>
    <footer>—Rob Pike, <cite>Go Proverbs - Rob Pike - Gopherfest - November 18, 2015<br> <a href="https://www.youtube.com/watch?v=PAAkCSZUG1c&t=2m48s">https://www.youtube.com/watch?v=PAAkCSZUG1c&t=2m48s</a></cite></footer>
</blockquote>
