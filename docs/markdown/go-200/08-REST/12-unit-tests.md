<!-- .slide: class="with-code" -->

# Tests HTTP

- Le package httptest
- Les tests unitaires avec httptest.ResponseRecorder

```go
func TestSomeHandler(t *testing.T) {
    handler := func(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "Hello World!")
    }
    req := httptest.NewRequest("GET", "http://example.com/foo", nil)
    w := httptest.NewRecorder()
    handler(w, req)
    resp := w.Result()
    fmt.Println(resp.StatusCode)
   //...
}
```

Notes:
OFU

ResponseRecorder respecte l’interface ResponseWriter et enregistre ce qui se passe
