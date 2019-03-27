FROM alpine:latest

MAINTAINER jianliu001922@gmail.com

ENV APP_SOURCE_CODE_PATH /opt/app
ENV APP_LOG_PATH /data/app/log/
ENV APP_NAME facade_server

WORKDIR ${APP_SOURCE_CODE_PATH}
VOLUME ${APP_LOG_PATH}

RUN echo "PS1=$" >> ~/.bashrc

ADD ${APP_NAME} .

CMD ${APP_NAME}
