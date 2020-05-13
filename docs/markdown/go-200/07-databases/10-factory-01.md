<!-- .slide: class="with-code" -->

# Accès aux données

## La fabrique - consts & vars

```go
// DBType lists the type of implementation the factory can return
type DBType intconst (
    // DAOMongo is used for Mongo implementation of TaskDAO
    DAOMongo DBType = iota
    // DAOMock is used for mocked implementation of TaskDAO
    DAOMock
    // mongo timeout
    timeout = 5 \* time.Second
    // poolSize of mongo connection pool
    poolSize = 35
)

var (
    // ErrorDAONotFound is used for unknown DAO type
    ErrorDAONotFound = errors.New("unknown DAO type")
)
```

Notes:
SFR

constantes publiques

erreur publiques
