.PHONY: spec serve ui

spec:
	cd cmd/serve/; swagger generate spec -o ../../swagger.json

swag:
	swagger serve https://raw.githubusercontent.com/djacobs24/go-swagger-example/master/swagger.json
