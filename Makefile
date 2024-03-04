build:
	@go build -o ./bin/parkingspace

run: build
	@./bin/parkingspace


test:
	@go test ./...