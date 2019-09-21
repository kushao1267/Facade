include .env

build-prod:
	GO111MODULE=on CGO_ENABLED=0 GOOS=$(GOOS) $(GO) build -o ./bin/$(APP_NAME) -mod=vendor main.go

build-image-prod: build-prod
	# 构建生产镜像
	docker build -t facade .

prod: build-image-prod
	# 启动整个项目,生产环境
	docker-compose up

prod-down: 
	# 停掉整个项目,生产环境
	docker-compose down

dev:
	# 运行开发环境(热启动)
	gin -a 8080 -p 3000

exec-prod:
	# 进入容器
	docker exec -it $(APP_NAME) sh

test:
	# 运行测试
	@GO111MODULE=on $(GO) test -mod=vendor -v ./facade/techniques/*

tidy:
	$GO111MODULE=on (GO) mod tidy
	$GO111MODULE=on (GO) mod vendor

lint:
	@golint

clean:
	@GO111MODULE=on $(GO) clean -mod=vendor && rm -rf ./bin
