run: generate-frontend
	go run main.go

run-frontend:
	cd ui && yarn dev

generate-frontend:
	cd ui && yarn generate && rm -rf ../static && cp -r dist ../static

serve-swagger:
	docker run -p 8000:8080 -e SWAGGER_JSON=/api.yaml -v ./docs/swagger.yaml:/api.yaml swaggerapi/swagger-ui

generate-api-docs:
	./swag init

build-mac: generate-frontend
	env GOOS=darwin GOARCH=amd64 go build -o bin/listings-darwin

build-linux: generate-frontend
	env GOOS=linux GOARCH=amd64 go build -o bin/listings-linux

install:
	cd ui && yarn install

open:
	open http://localhost:8080

open-swagger:
	open http://localhost:8000