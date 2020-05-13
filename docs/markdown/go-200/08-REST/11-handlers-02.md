<!-- .slide: class="with-code-bg-dark" -->

# REST

## Fonctions handler associées aux routes

Implémentation des `HandlerFunc func(ResponseWriter, *Request)`

pour les utilitaires JSON, Go a ce qu’il faut : voir _web/utils.go_

```go
// Get retrieve an entity by id
func (tc *TaskController) Get(w http.ResponseWriter, r *http.Request) {
    // get the task's ID from the URL
    taskID := ParamAsString("id", r)
    // [...]
    logger.WithField("tasks", task).Debug("task found")
    SendJSONOk(w, task)
}
```
