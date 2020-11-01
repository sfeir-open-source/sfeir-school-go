package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

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
			SetSocketTimeout(timeout))
		// TODO set the connection pool size using the poolSize const

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

		// TODO return a new DAO Mongo build with the configured session
		return nil, nil
	case DAOPostgres:
		fallthrough
	case DAOMock:
		return NewTaskDAOMock(), nil
	default:
		return nil, ErrorDAONotFound
	}
}
