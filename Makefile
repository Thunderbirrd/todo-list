swag:
	swag init -g cmd/main.go

build-image:
	docker build -t todo-list .

docker-start:
	docker run --name=todo-list --publish 8080:8080 --rm todo-list