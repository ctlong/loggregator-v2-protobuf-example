build-old-server:
	go build -o bin/old-server cmd/old/server/main.go

build-old-client:
	go build -o bin/old-client cmd/old/client/main.go

build-new-server:
	go build -o bin/new-server cmd/new/server/main.go

build-new-client:
	go build -o bin/new-client cmd/new/client/main.go

run-old-server: build-old-server
	./bin/old-server

run-old-client: build-old-client
	./bin/old-client

run-new-server: build-new-server
	./bin/new-server

run-new-client: build-new-client
	./bin/new-client

build: build-old-server build-old-client build-new-server build-new-client
