include .env

.PHONY: build-example build run exec

build-example:
	CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} ${GO} build -a -installsuffix cgo -o ${APP_NAME} ./examples/main.go

build:
	# 构建项目镜像
	docker build -t facade .

run:
	# 运行:
	docker run --rm -it -p 8080:8080 --name my-facade facade

exec:
	# 进入容器:
	docker exec -it my-facade /bin/bash
