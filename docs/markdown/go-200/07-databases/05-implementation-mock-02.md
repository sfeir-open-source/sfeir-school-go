<!-- .slide: class="with-code-bg-dark" -->

# Accès aux données

## Implémentation : Mock

```go
// MockedTask is the task returned by this mocked interface
var MockedTask = model.Task{
    ID:             uuid.New().String(),
    Title:          "Learn Go",
    Description:    "Let's learn the Go programming language.",
    Status:         model.StatusInProgress,
    Priority:       model.PriorityHigh,
    CreationDate:   time.Date(2017, 01, 01, 0, 0, 0, 0, time.UTC),
    DueDate:        time.Date(2017, 03, 23, 0, 0, 0, 0, time.UTC),
}

// Compile time check
var _ TaskDAO = (*TaskDAOMock)(nil)
```

Notes:
OFU

Implémentation d’une interface en respectant les méthodes et leurs signatures
