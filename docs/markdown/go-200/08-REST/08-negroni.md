<!-- .slide: class="with-code-bg-dark" -->

# REST

## Negroni Middleware, de quoi parle-t-on ?

Interface Handler similaire au Handler http MAIS avec un param next

```go
// Handler handler is an interface that objects can implement to be registered to serve as middleware in the Negroni middleware stack.
type Handler interface {
    ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

// HandlerFunc is an adapter to allow the use of ordinary functions as Negroni handlers.
type HandlerFunc func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)

func (h HandlerFunc) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    h(rw, r, next)
}
```

![center](./assets/go-200/images/FrancoisPignon.png)

Notes:
SFR

idem node, spring filter, etc...
