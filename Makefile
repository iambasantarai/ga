build: 
	go build -o bin/ga
test: build
	go test ./... -count=1
token: build
	./bin/ga ./examples/token.ga
