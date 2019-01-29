.PHONY: spec serve ui

spec:
	cd cmd/serve/; swagger generate spec -o ../../swagger.json

serve:
	go run ./cmd/serve/main.go

ui:
	docker run --rm -it -p 8081:8080 swaggerapi/swagger-ui
