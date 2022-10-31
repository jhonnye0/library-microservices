HELLO_IMAGE = api-gateway:latest

run/hello:
	go run hello/main.go

docker/run/hello:
	docker run -p 8080:8080 $(HELLO_IMAGE)

tidy:
	go mod tidy

build/hello:
	docker build -t $(HELLO_IMAGE) .