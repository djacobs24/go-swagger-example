# Go Swagger Example

This is an example of [go-swagger](https://github.com/go-swagger/go-swagger).

`go-swagger` is used to generate swagger specifications for API documentation.

# Dep
Go Swagger Example uses dep to manage dependencies.

`dep ensure` - syncs project dependencies

`dep ensure -update` - updates all dependencies

# Commands
Go Swagger Example inclues a [makefile](https://github.com/djacobs24/go-swagger-example/blob/master/Makefile)

`make spec` - generates the Swagger JSON spec

`make swag` - serves swagger ui for the spec