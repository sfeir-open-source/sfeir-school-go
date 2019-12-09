# SFEIR School Go

This project is meant to help learning go. It provides a basic implementation of a REST microservice exposing a CRUD API.
Data are persisted in a MongoDB NoSQL database and the application is deployed in Docker.

## Technical stack

* [Go is the language](https://golang.org)
* [Docker](https://www.docker.com)
* [MongoDB NoSQL database](https://www.mongodb.com)
* [The mongodb driver](https://go.mongodb.org/mongo-driver)
* [PostgreSQL database](https://www.postgresql.org)
* [PostgreSQL Go driver](https://github.com/lib/pq)
* [Database migration](https://github.com/golang-migrate/migrate)
* [Gorilla Mux the URL router](https://github.com/gorilla/mux)
* [Urfave negroni Web HTTP middleware](https://github.com/urfave/negroni)
* [Urfave cli the command line client parser](https://github.com/urfave/cli/v2)
* [Sirupsen the logger](https://github.com/sirupsen/logrus)
* [Golint the source linter](https://github.com/golang/lint)

## Architecture

![main architecture](docs/img/main_architecture.png "Main architecture")

![web architecture](docs/img/web_architecture.png "Web architecture")

## Build

```shell
make help
```
