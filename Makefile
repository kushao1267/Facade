include .env

.PHONY: build-image compose build exec tidy test clean

build-image: build
	# 构建项目镜像
	docker build -t facade .

compose: build
	# 构建项目镜像
	docker-compose up

build:
	GOOS=$(GOOS) $(GO) build -o ./bin/$(APP_NAME) -mod=vendor main.go

exec:
	# 进入容器:
	docker exec -it $(APP_NAME) sh

test:
	# 运行测试
	@$(GO) test -mod=vendor -v ./facade/techniques/*

tidy:
	$(GO) mod tidy
	$(GO) mod vendor

clean:
	@$(GO) clean -mod=vendor && rm -rf ./bin
