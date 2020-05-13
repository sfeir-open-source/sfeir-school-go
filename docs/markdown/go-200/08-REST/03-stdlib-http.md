<!-- .slide: class="with-code-bg-dark" -->

# REST

## Go arrive avec une bibliothèque web bien fournie : le package http

En quelques lignes vous pouvez monter un serveur web avec :

- support HTTP/1 et HTTP/2 (Go 1.6), Push (Go 1.8)
- HTTPS

```go
func main() {
    http.Handle("/", http.FileServer(http.Dir("./")))
    http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusNotImplemented) // 501
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        fmt.Fprintln(w, "Not implemented yet")
    }
)

http.ListenAndServe(":8080", nil)}
```

![center](./assets/go-200/images/magic.webp)

<!-- .element style="position:absolute; right: 50px; top: 50px" -->

Notes:
SFR

7 lignes, un serveur de fichiers web + un end point en plus.

magie, si on creuse
