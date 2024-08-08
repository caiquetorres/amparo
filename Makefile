build:
	@go build -o bin/amparo cmd/main.go

run: build
	@./bin/amparo

test:
	@go test ./... -v

deploy:
	@GOARCH=arm64 GOOS=linux go build -o bin/bootstrap cmd/main.go
	@zip -rj bootstrap.zip bin/bootstrap
