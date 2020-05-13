<!-- .slide: class="with-code" -->

# Parsing des arguments

## run

```shell
$ go run main.go -port 8090

port : 8090

$ go run run main.go -port 8090xy

Incorrect Usage. invalid value "8090xy" for flag -port: strconv.ParseInt: parsing "8090xy": invalid syntax

NAME:
mytodolist - mytodolist service launcher

[…]

GLOBAL OPTIONS:
--port value Set the listening port of the webserver (default: 8020)
--help, -h show help
--version, -v print the version
```
