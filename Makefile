include .env

<<<<<<< HEAD
GO=$(go env GOROOT)/bin/go
=======
build-prod:
	GO111MODULE=on CGO_ENABLED=0 GOOS=$(GOOS) $(GO) build -o ./bin/$(APP_NAME) -mod=vendor main.go
>>>>>>> 41768819d1a05f2607694f2f4665b9db54b88748

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
<<<<<<< HEAD
	$(GO) mod tidy
=======
	$GO111MODULE=on (GO) mod tidy
	$GO111MODULE=on (GO) mod vendor
>>>>>>> 41768819d1a05f2607694f2f4665b9db54b88748

.PHONY:lint
lint:
	@golint

<<<<<<< HEAD
=======
clean:
	@GO111MODULE=on $(GO) clean -mod=vendor && rm -rf ./bin
>>>>>>> 41768819d1a05f2607694f2f4665b9db54b88748
