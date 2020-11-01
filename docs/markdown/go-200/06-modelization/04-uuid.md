<!-- .slide: class="with-code" -->

# Modélisation

## Gestion de l’ID unique par un UUID

```go
import (
    "github.com/google/uuid"
    "time"
)

// NewTask sets a new ID of the Task as a string
func NewTask(title, description string) *Task {
    return &Task{
        ID:             uuid.New().String(),
        CreationDate:   time.Now(),
        DueDate:        time.Now().Add(24 * time.Hour),
        Status:         StatusTodo,
        Priority:       PriorityMedium,
        Title:          Title,
        Description     Description,
    }
}
```
