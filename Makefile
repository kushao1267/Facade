include .env

.PHONY: build-example build run exec tidy test

build:
	# 构建项目镜像
	docker build -t facade .

build-example:
	GOOS=${GOOS} ${GO} build -o ${APP_NAME} examples/main.go
	docker run --rm -it -p 8080:8080 --name ${APP_NAME} -e "APP_ENV=${APP_ENV}" -v "${PWD}":/go/ facade

run:
	# 运行:
	docker run --rm -it -p 8080:8080 --name ${APP_NAME} -e "APP_ENV=${APP_ENV}" -v "${PWD}":/go/ facade

exec:
	# 进入容器:
	docker exec -it ${APP_NAME} /bin/bash

test:
	# run test

tidy:
	go mod tidy