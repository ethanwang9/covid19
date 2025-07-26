FROM golang:1.19-alpine as builder

WORKDIR /go/src/github.com/ethanwang9/covid19/server
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="Ethan.Wang@social@ethan9.cn"

WORKDIR /app

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk --no-cache add tzdata  && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

COPY --from=builder /go/src/github.com/ethanwang9/covid19/server/server ./
COPY --from=builder /go/src/github.com/ethanwang9/covid19/server/config/config.ini ./config/config.ini
COPY --from=builder /go/src/github.com/ethanwang9/covid19/server/config/c19.sql ./config/c19.sql

VOLUME /app/log

EXPOSE 8088
ENTRYPOINT ./server
