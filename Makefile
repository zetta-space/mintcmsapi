build:
	@go build -o bin/mintcmsapi

run: build
	@./bin/mintcmsapi

test: 
	@go test ./... -v