FROM alpine:3.6

ENV app_name facade_server
ENV app_dir /opt/app

WORKDIR ${app_dir}

RUN echo http://mirrors.aliyun.com/alpine/v3.6/main > /etc/apk/repositories \
    && echo http://mirrors.aliyun.com/alpine/v3.6/main >> /etc/apk/repositories \
    && apk add --no-cache make \
    && apk add --no-cache ca-certificates

CMD ./bin/${app_name}