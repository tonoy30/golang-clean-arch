all: clean build run
clean:
	rm -rf ./bin
build:
	go build -o bin/main cmd/main.go
run:
	./bin/main
