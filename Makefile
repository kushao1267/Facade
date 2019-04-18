include .env

.PHONY: build-example build run exec tidy test

run-example: build-example run

build:
	# 构建项目镜像
	docker build -t facade .

build-example:
	GOOS=${GOOS} ${GO} build -o ${APP_NAME} -mod=vendor examples/main.go

run:
	# 运行:
	docker run --rm -it -p 8080:8080 --name ${APP_NAME} -e "APP_ENV=${APP_ENV}" -v "${PWD}":/go/ facade

exec:
	# 进入容器:
	docker exec -it ${APP_NAME} /bin/bash

test:
	# 运行测试
	go test -v -mod=vendor ./facade/techniques/*

tidy:
	go mod tidy
