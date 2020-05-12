<!-- .slide: class="with-code-bg-dark" -->

# Accès aux données

## Implémentation : Mongodb

```go
// GetByID returns a task by its ID
func (s *TaskDAOMongo) GetByID(ID string) (model.Task, error) { // check ID
    if _, err := uuid.Parse(ID); err != nil {
        return nil, ErrInvalidUUID
    }

    // use client to retrieve single result
    task := model.Task{}
    c := s.client.Database(todolist).Collection(tasks)
    result := c.FindOne(context.TODO(), bson.M{"id": ID})

    // decode result and handle error
    err = result.Decode(&task)

    return &task, err
}
```

Notes:
SFR

ORM possible
