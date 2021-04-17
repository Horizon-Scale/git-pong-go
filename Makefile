fmt:
	@go fmt ./... | gofmt -s -w .

lint:
	@go vet ./...

tests:
	@go test ./...

start:
	@go run ./cmd/pong-api

build:
	@go build -tags netgo -ldflags="-s -w" -trimpath -o dist/pong-api ./cmd/pong-api

clean:
	@rm -rf dist

update:
	@go mod tidy
