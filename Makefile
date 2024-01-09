
build-writer:
	go build -o bin/writer cmd/writer/main.go

build-playground:
	go build -o bin/playground cmd/playground/main.go

build-server:
	go build -o bin/server cmd/server/main.go


build: build-writer build-playground build-server