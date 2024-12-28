build: 
	go build -o bin/ga
test: build
	go test ./... -count=1
example: build
	./bin/ga ./examples/example.ga
