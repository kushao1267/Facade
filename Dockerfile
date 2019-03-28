FROM golang:1.12.1-stretch

MAINTAINER jianliu001922@gmail.com

ENV APP_LOG_PATH /data/app/log/
ENV APP_NAME facade_server

WORKDIR /go/
VOLUME ${APP_LOG_PATH}

RUN echo "PS1=$" >> ~/.bashrc
COPY . .

CMD ./${APP_NAME}