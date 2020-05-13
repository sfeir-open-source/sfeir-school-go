<!-- .slide: class="with-code-bg-dark" -->

# REST

## Handler, HandlerFunc : Interface et fonctions anonymes qui répondent à des Requêtes.

```go
// A Handler responds to an HTTP request.
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

// The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers.
// If f is a function with the appropriate signature,
// then HandlerFunc(f) is a Handler that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
```
