<!-- .slide: class="with-code" -->

# Accès aux données

## Implémentation : Mock

```go
// TaskDAOMock is the mocked implementation of the TaskDAO
type TaskDAOMock struct {
    storage map[string]model.Task
}

// NewTaskDAOMock creates a new TaskDAO with a mocked implementation
func NewTaskDAOMock() TaskDAO {
    daoMock := &TaskDAOMock{
        storage: make(map[string]model.Task),
    }
    // Adds some fake data
    daoMock.Save(MockedTask)
    return daoMock
}
```

Notes:
OFU

getBy ⇒ passage d’une fonction en paramètre
