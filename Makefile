include .env

.PHONY: build-image compose build run exec tidy test dev clean

build-image: build
	# 构建项目镜像
	docker build -t facade .

compose: build
	# 构建项目镜像
	docker-compose up

build:
	GOOS=$(GOOS) $(GO) build -o ./bin/$(APP_NAME) main.go

run:
	# 运行:
	docker run --rm -it -p 8080:8080 \
		--name $(APP_NAME) \
		-e "GIN_MODE=$(GIN_MODE)" \
		-v "${PWD}/bin":$(WORK_DIR)/bin facade
	

dev: build run

exec:
	# 进入容器:
	docker exec -it $(APP_NAME) sh

test:
	# 运行测试
	@go test -v ./facade/techniques/*

tidy:
	go mod tidy
	go mod vendor

clean:
	@ go clean && rm -rf ./bin
