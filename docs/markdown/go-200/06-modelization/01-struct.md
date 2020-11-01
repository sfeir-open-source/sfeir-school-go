<!-- .slide: class="with-code" -->

# Modélisation

Première étape pour la couche d’accès aux données : la modélisation

Une structure de base définissant ce qu’est une Task :

```go
type Task struct {
    ID              string
    Title           string
    Description     string
    Status          int // Todo, In Progress, Done
    Priority        int // High, Medium, Low
    CreationDate    time.Time
    DueDate         time.Time // Echéance
}
```

<!-- .element: class="big-code" -->
