<!-- .slide: class="with-code-bg-dark" -->

# Accès aux données

## La fabrique - Mock & Mongodb

```go
// GetTaskDAO returns a TaskDAO according to type and params
func GetTaskDAO(param string, daoType DBType) (TaskDAO, error) {
    switch daoType {
        case DAOMock:
            return NewTaskDAOMock(), nil
        case DAOMongo:
            ctx, cancel := context.WithTimeout(context.Background(), timeout)
            defer cancel()
            client, err := mongo.Connect(
                ctx,
                options.Client().
                    .ApplyURI(cnxStr)
                    .SetSocketTimeout(timeout)
                    .SetMaxPoolSize(poolSize)
            )
    	    // add ping for database availability
            return NewTaskDAOMongo(client), nil
    }
}
```
