<!-- .slide: class="with-code" -->

# Accès aux données

## La fabrique - Postgresql (suite)

```go
    // ...

    // check is db migration is necessary
    if len(migrationPath) == 0 {
        return NewTaskDAOPostgres(db), nil
    }

    //  playing database migration
    driver, err := postgres.WithInstance(db, &postgres.Config{})
    m, err := migrate.NewWithDatabaseInstance(
        fileScheme+migrationPath,
        "postgres", driver)

    if err != nil {
        return nil, err
    }

    // upgrade database if necessary
    err = m.Up()
    if err != nil {
        if err != migrate.ErrNoChange {
            return nil, err
        }
    }
    // ...
}
```

<!-- .element style="font-size:0.6em; line-height: 1em" -->
