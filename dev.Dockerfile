FROM golang:1.19-alpine as builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && \
    apk --no-cache add git
COPY . /app
WORKDIR /app
RUN mkdir -p target && \
    GO111MODULE=on GOOS=linux GOPROXY=https://goproxy.cn,direct go build -o target/ender_chest main.go


FROM alpine:3.19 as prod

WORKDIR /app
COPY --from=0 /app/target/ender_chest .

ENTRYPOINT ["/app/ender_chest"]