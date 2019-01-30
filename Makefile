.PHONY: clean spec valid swag

clean:
	rm ./swagger.json

spec:
	cd cmd/serve/; swagger generate spec --scan-models -o ../../swagger.json

valid:
	swagger validate ./swagger.json

swag:
	swagger serve ./swagger.json
