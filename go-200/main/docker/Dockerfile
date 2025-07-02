# build stage
FROM golang:1.15-alpine AS build-env

RUN apk --no-cache add -U make git

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .

RUN make all

# final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=build-env /app/todolist /app/

# by default, the exposed ports are 8020 (HTTP)
EXPOSE 8020
ENTRYPOINT ["/app/todolist"]
