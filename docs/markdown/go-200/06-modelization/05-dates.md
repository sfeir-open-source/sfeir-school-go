<!-- .slide: class="with-code" -->

# Modélisation

## Comparaison avec Equal à cause des dates

Pour comparer 2 structures en Go toutes ses propriétés doivent être comparables

```go
// Equal compares a Task to another
func (t Task) Equal(task Task) bool {
    return t.ID == task.ID &&
        t.Title == task.Title &&
        t.Description == task.Description &&
        t.Status == task.Status &&
        t.Priority == task.Priority &&
        t.CreationDate.Equal(task.CreationDate) &&
        t.DueDate.Equal(task.DueDate)
}
```

<!-- .element: class="big-code" -->

Explication pour les dates : https://golang.org/pkg/time/#Time
