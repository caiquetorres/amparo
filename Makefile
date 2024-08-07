build:
	@go build -o bin/amparo cmd/main.go

run: build
	@./bin/amparo

test:
	@go test ./... -v
