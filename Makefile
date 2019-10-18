include .env

GO=$(go env GOROOT)/bin/go

.PHONY:build-image-prod
build-image-prod:
	# 构建生产镜像
	docker build --build-arg APP_NAME=Facede \
				 -f Dockerfile . \
				 -t "kushao1267/facede:latest"

.PHONY:up
up: build-image-prod
	# 启动整个项目,生产环境
	docker-compose up

.PHONY:down
down:
	# 停掉整个项目,生产环境
	docker-compose down

.PHONY:dev
dev:
	# 运行开发环境(热启动)
	gin -a 8080 -p 3000

.PHONY:exec-prod
exec-prod:
	# 进入容器
	docker exec -it $(APP_NAME) sh

.PHONY:test
test:
	# 运行测试
	@GO111MODULE=on $(GO) test -mod=vendor -v ./facade/techniques/*

.PHONY:tidy
tidy:
	$(GO) mod tidy

.PHONY:lint
lint:
	@golint

