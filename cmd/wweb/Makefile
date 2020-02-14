.PHONY: *

CONTAINER_NAME=wweb

build:
	go build -o -v
test:
	go test -v ./...
run:
	go run *.go run

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v
docker-clean:
	docker rm $(CONTAINER_NAME)
	docker rmi -f $(CONTAINER_NAME) 
docker-run: Dockerfile build-linux
	docker build -t $(CONTAINER_NAME) .
	docker run -d -p 8888:8080 --name "$(CONTAINER_NAME)" $(CONTAINER_NAME)
docker-stop: 
	docker  stop "$(CONTAINER_NAME)" 
