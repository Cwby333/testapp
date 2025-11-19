build_run: build run

run:
	sudo docker compose up 

build:
	GOOS=linux GOARCH=amd64 go build -o main ./internal/cmd/testapp/main.go
	sudo docker build --platform linux/amd64 -t testapp:1.0 .