# Makefile for Todolist : GO-200

.DEFAULT_GOAL := all

# -----------------------------------------------------------------
#        Version
# -----------------------------------------------------------------

# version
VERSION=2.0.0
BUILDDATE=$(shell date -u '+%s')
BUILDHASH=$(shell git rev-parse --short HEAD)
VERSION_FLAG=-ldflags "-X main.Version=$(VERSION) -X main.GitHash=$(BUILDHASH) -X main.BuildStmp=$(BUILDDATE)"

# -----------------------------------------------------------------
#        Main targets
# -----------------------------------------------------------------

all: clean build ## Clean and build the project

clean: ## Clean the project
	@go clean -r -cache -testcache ./...
	@go mod tidy
	@rm -Rf .tmp .DS_Store *.log *.out *.mem *.test build/

build: format ## Build all libraries and binaries
	@CGO_ENABLED=0 go build -v $(VERSION_FLAG) -o todolist todolist.go

format: ## Format all packages
	@go fmt ./...

teardownTest: ## Tear down databases for integration tests
	@docker-compose -f docker/databases.yml down

setupTest: teardownTest ## Start databases for integration tests
	@docker-compose -f docker/databases.yml up -d

test: setupTest ## Start tests with a databases docker image
	@export DB_HOST=localhost; sleep 2; go test -v ./...; make teardownTest

bench: setupTest ## Start benchmark
	@go test -v -run nothing -bench=. -memprofile=prof.mem github.com/sfeir-open-source/sfeir-school-go/web ; make teardownTest

benchTool: bench ## Start benchmark tool
	@echo "### TIP : type 'top 5' and 'list path_of_the_first_item'"
	@go tool pprof --alloc_space web.test prof.mem

lint: ## Lint all packages
	@golint ./...
	@go vet ./...

# -----------------------------------------------------------------
#        Docker targets
# -----------------------------------------------------------------

dockerBuild: ## Build a docker image of the program
	docker build -f docker/Dockerfile -t sfeir/todolist:latest .

dockerClean: ## Remove the docker image of the program
	docker rmi -f sfeir/todolist:latest

dockerUp: ## Start the program instances with their databases
	docker-compose -f docker/docker-compose.yml up -d

dockerDown: ## Stop the program instances, their databases and remove the containers
	docker-compose -f docker/docker-compose.yml down

dockerBuildUp: dockerDown dockerBuild dockerUp ## Stop, build and launch the docker images of the program

dockerWatch: ## Watch the status of the docker container
	@watch -n1 'docker ps | grep todolist'

dockerLogs: ## Print the logs of the container
	docker-compose logs -f

help: ## Print this message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: all test clean teardownTest setupTest
