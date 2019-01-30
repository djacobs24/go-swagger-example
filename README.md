# Go Swagger Example

This is an example of [go-swagger](https://github.com/go-swagger/go-swagger).

`go-swagger` is used to generate swagger specifications for API documentation.

# Commands
Go Swagger Example inclues a [makefile](https://github.com/djacobs24/go-swagger-example/blob/master/Makefile)

`make clean` - removes the Swagger JSON specification

`make spec` - generates the Swagger JSON specification

`make valid` - validates the Swagger JSON specification

`make swag` - serves Swagger UI for the specification

# Instructions
If annotations are modified run `make clean` to remove the Swagger JSON spec, `make spec` to regenerate the spec, and `make valid` to validate the spec. 

If the spec is valid, or no changes were made, run `make swag` to serve the API documentation.