<!-- .slide: class="with-code-bg-dark" -->

# Accès aux données

## Implémentation : Mongodb

```go
const (
    todolist = "todolist"
    tasks = "tasks"
    index = "id"
)

// TaskDAOMongo is the mongo implementation of the TaskDAO
type TaskDAOMongo struct {
    client *mongo.Client
}

// NewTaskDAOMongo creates a new TaskDAO mongo implementation
func NewTaskDAOMongo(client *mongo.Client) TaskDAO {
    return &TaskDAOMongo{
        client: client,
    }
}
```
