.PHONY: spec valid swag

spec:
	cd cmd/serve/; swagger generate spec --scan-models -o ../../swagger.json

valid:
	swagger validate ./swagger.json

swag:
	swagger serve ./swagger.json
