<!-- .slide: class="with-code-bg-dark" -->

# Accès aux données

## Implémentation : Postgresql

```go
// TaskDAOPostgres is the postgres implementation of the TaskDAO
type TaskDAOPostgres struct {
    db *sql.DB
}

// NewTaskDAOPostgres creates a new TaskDAO postgres implementation
func NewTaskDAOPostgres(db *sql.DB) TaskDAO {
    return &TaskDAOPostgres{
        db: db,
    }
}
```
