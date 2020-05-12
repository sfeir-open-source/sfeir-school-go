<!-- .slide: class="with-code-bg-dark" -->

# Accès aux données

## Implémentation : Postgresql

```go
// GetByID returns a task by its ID
func (s *TaskDAOPostgres) GetByID(ID string) (model.Task, error) {

    // check ID...

    // query db
    rows, err := s.db.Query("SELECT * FROM todos WHERE uuid=$1", ID)
    if err != nil {
    	return nil, err
    }

    results, err := mapRows(rows)

    if len(results) == 0 {
    	return nil, ErrNotFound
    }

    //[...]
}
```

Notes:
SFR

Pas d’ORM avancé, pas la tendance en Go
