include .env

.PHONY:build-image-prod compose-prod compose-prod-down build-prod dev prod exec-prod test tidy lint clean

build-prod:
	CGO_ENABLED=0 GOOS=$(GOOS) $(GO) build -o ./bin/$(APP_NAME) -mod=vendor main.go

build-image-prod: build-prod
	# 构建生产镜像
	docker build -t facade .

compose-prod: build-image-prod
	# 启动整个项目,生产环境
	docker-compose up

compose-prod-down: build-image-prod
	# 启动整个项目,生产环境
	docker-compose down

dev:
	# 运行开发环境
	@GIN_MODE=test gowatch -o ./bin/facade_dev_server -p -mod=vendor .

prod:
	# 运行正式环境
	./bin/facade_server

exec-prod:
	# 进入容器
	docker exec -it $(APP_NAME) sh

test:
	# 运行测试
	@$(GO) test -mod=vendor -v ./facade/techniques/*

tidy:
	$(GO) mod tidy
	$(GO) mod vendor

lint:
	@golint

clean:
	@$(GO) clean -mod=vendor && rm -rf ./bin && rm -f gin-bin
