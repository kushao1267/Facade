FROM alpine:3.6

ENV app_name facade_server
ENV app_dir /opt/app/

ARG port=8080

RUN mkdir -p app_dir
WORKDIR ${app_dir}

COPY Makefile .
COPY config.toml .
COPY .env .

RUN echo http://mirrors.aliyun.com/alpine/v3.6/main > /etc/apk/repositories \
    && echo http://mirrors.aliyun.com/alpine/v3.6/main >> /etc/apk/repositories \
    && apk add --no-cache make \
    && apk add --no-cache ca-certificates

EXPOSE $port

CMD ./bin/${app_name}