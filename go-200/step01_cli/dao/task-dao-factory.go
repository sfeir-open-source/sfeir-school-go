package dao

import (
	"context"
	"database/sql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	// importing postgres driver for sql connection
	_ "github.com/lib/pq"
	// importing file source for db migration
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"time"
)

const (
	// db timeout
	timeout = 5 * time.Second

	// poolSize of db connection pool
	poolSize = 35

	// file url scheme
	fileScheme = "file://"
)

// GetTaskDAO returns a TaskDAO according to type and params
func GetTaskDAO(cnxStr, migrationPath string, daoType DBType) (TaskDAO, error) {
	switch daoType {
	case DAOMongo:

		// mongo connection

		// connect with params
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().
			ApplyURI(cnxStr).
			SetSocketTimeout(timeout).
			SetMaxPoolSize(poolSize))

		// check error
		if err != nil {
			return nil, err
		}

		// ping to ensure connection is ok
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
		defer cancel()
		if err = client.Ping(ctx, readpref.Primary()); err != nil {
			return nil, err
		}

		return NewTaskDAOMongo(client), nil
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

		return NewTaskDAOPostgres(db), nil
	case DAOMock:
		return NewTaskDAOMock(), nil
	default:
		return nil, ErrorDAONotFound
	}
}
