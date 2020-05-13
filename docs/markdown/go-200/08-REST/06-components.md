# REST

## routeur, middleware et controller

- Pour faire du web il nous faut :

  - Un serveur web
  - Un routeur
  - Des HandleFunc en face de chaque route
  - En bonus des middleware

- Pour une utilisation un peu plus fun, on va utiliser les bibliothèques suivantes compatible avec les standards du package http :
  - https://github.com/urfave/negroni : pour la gestion des middlewares
  - https://github.com/gorilla/mux : un routeur plus “dev friendly”

NB. Il existe aussi [gin](https://github.com/gin-gonic/gin) et [martini](https://github.com/go-martini/martini) mais plus complexes et sans http Go (voire vraiment magiques)

![center h-300](./assets/go-200/images/cocktail.webp)

<!-- .element style="position:absolute; right: 50px; top: 50px" -->
