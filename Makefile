build: 
	@go build -o bin/pipes
	 
run: build
	@./bin/pipes

test: 
	@go test -v ./...