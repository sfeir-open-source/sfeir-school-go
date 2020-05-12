<!-- .slide: class="with-code-bg-dark" -->

# Modélisation

## Définir des “énumérés” type safe pour les constantes

```go
// TaskPriority is the priority of a task
type TaskPriority int
const (
    // PriorityMinor lower priority
    PriorityMinor TaskPriority = iota
    // PriorityMedium medium priority
    PriorityMedium
    // PriorityHigh high priority
    PriorityHigh
)

// TaskStatus is the status of a task
type TaskStatus int
const (
    // StatusTodo for incomplete tasks
    StatusTodo TaskStatus = iota
    // StatusInProgress for tasks in progress
    StatusInProgress
    // StatusDone for completed tasks
    StatusDone
)
```
