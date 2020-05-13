<!-- .slide: class="with-code" -->

# REST

## Handle, HandleFunc : Fonctions qui travaillent sur une instance de routeur par défaut.

```go
// Handle registers the handler for the given pattern in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
func Handle(pattern string, handler Handler) {
    DefaultServeMux.Handle(pattern, handler)
}

// HandleFunc registers the handler function for the given pattern in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    DefaultServeMux.HandleFunc(pattern, handler)
}
```

<!-- .element: class="big-code" -->
