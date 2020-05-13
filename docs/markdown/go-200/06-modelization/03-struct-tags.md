<!-- .slide: class="with-code" -->

# Modélisation

## Ajouter des “**annotations**” pour les transformations **BSON** (Mongo) et **JSON** (Web REST)

Ce qui nous donne la structure finale suivante :

```go
type Task struct {
    ID              string          `json:"id,omitempty" bson:"id"`
    Title           string          `json:"title" bson:"title"`
    Description     string          `json:"description" bson:"description"`
    Status          TaskStatus      `json:"status" bson:"status"`
    Priority        TaskPriority    `json:"priority" bson:"priority"`
    CreationDate    time.Time       `json:"creation_date" bson:"creation_date"`
    DueDate         time.Time       `json:"due_date" bson:"due_date"`
}
```

<!-- .element: class="big-code" -->

💡Pro tip : Go 1.8 permet le cast de structures identiques ex: `taskDTO = TaskDTO(taskDAO)`

Notes:
OFU

cast go 1.8 entre DAO et DTO ! Au lieu d’un mapper

découplage BDD et Web, versionning
