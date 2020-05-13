<!-- .slide: class="with-code" -->

# REST

## Comment composer des middleware : web/web-server.go

```go
// middleware builder
n := negroni.New()
// add middleware for logging (negroni.Handler)
n.Use(negronilogrus.NewMiddlewareFromLogger(logger.StandardLogger(), "task"))
// add statistics middleware
n.Use(NewStatisticsMiddleware(statisticsDuration))

// add other middlewares as needed

// new router
router := web.NewRouter(web.NewTaskController())
// route handler goes last (http.Handler)
n.UseHandler(router)

// run web server
n.Run(":" + strconv.Itoa(port))
```

Notes:
SFR

CORS
