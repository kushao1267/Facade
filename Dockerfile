FROM golang:latest as builder

ENV GO111MODULE=on
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/
ENV SOURCE_PATH=github.com/kushao1267/Facade

ARG APP_NAME

COPY . /go/src/${SOURCE_PATH}
RUN cd /go/src/${SOURCE_PATH} \
    && CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOARM=6 \
    go build -ldflags \
    "-X 'main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'`' -X 'main.githash=`git describe --always`' -X 'main.goversion=`go version`'" \
    -o app main.go

FROM scratch

ARG APP_NAME
ENV SOURCE_PATH=github.com/kushao1267/Facade

COPY --from=builder /go/src/${SOURCE_PATH}/config.toml /opt/
COPY --from=builder /go/src/${SOURCE_PATH}/app /opt/

WORKDIR /opt/

ENTRYPOINT ["./app"]