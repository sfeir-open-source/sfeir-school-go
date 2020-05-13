<!-- .slide: class="with-code-bg-dark" -->

# REST

## Fonctions handler associées aux routes

Allons voir _web/controller.go_ et _web/router.go_

Déclaration des end points

```go
// Get
routes = append(routes, Route{
    Name: "Get one task",
    Method: http.MethodGet,
    Pattern: "/{id}",
    HandlerFunc: controller.Get,
})
```
