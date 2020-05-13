<!-- .slide: class="with-code" -->

# Accès aux données

## La fabrique - Postgresql

```go
    case DAOPostgres:
    	// postgresql connection
    	db, err := sql.Open("postgres", cnxStr)

    	// check errors
    	if err != nil {
    		return nil, err
    	}

    	// set max connection in pool
    	db.SetMaxOpenConns(poolSize)

    	// try to ping host
    	if err = db.Ping(); err != nil {
    		return nil, err
    	}

        // ...
    }
}
```
