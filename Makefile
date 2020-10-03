clean:
	rm -rf bin/*

build:
	go build -o bin/main.exe

execute:
	./bin/main.exe

run: build execute

test:
	go test -cover -race ./...