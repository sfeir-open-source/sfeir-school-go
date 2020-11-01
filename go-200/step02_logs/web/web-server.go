package web

import (
	"github.com/meatballhat/negroni-logrus"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/rs/cors"
	"github.com/sfeir-open-source/sfeir-school-go/dao"
	logger "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"time"
)

// BuildWebServer constructs a new web server with the right DAO and tasks handler
func BuildWebServer(db, migration string, daoType dao.DBType, statisticsDuration time.Duration) (*negroni.Negroni, error) {

	// task dao
	td, err := dao.GetTaskDAO(db, migration, daoType)
	if err != nil {
		logger.WithField("error", err).WithField("dbtype", daoType).WithField("params", db).
			Warn("unable to build the required DAO")
		return nil, err
	}

	// web server
	n := negroni.New()

	// add middleware for logging
	n.Use(negronilogrus.NewMiddlewareFromLogger(logger.StandardLogger(), "task"))

	// add recovery middleware in case of panic in handler func
	recovery := negroni.NewRecovery()
	recovery.PrintStack = false
	n.Use(recovery)

	// add statistics middleware
	n.Use(NewStatisticsMiddleware(statisticsDuration))

	// add CORS (all origins, all methods)
	n.Use(cors.AllowAll())

	// add Gzip middleware
	n.Use(gzip.Gzip(gzip.DefaultCompression))

	// new controller
	controller := NewTaskController(td)

	// new router
	router := NewRouter(controller)

	// route handler goes last
	n.UseHandler(router)

	return n, nil
}
