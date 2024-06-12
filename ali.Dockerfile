FROM golang:1.19-alpine as builder

RUN apk --no-cache add git
COPY . /app
WORKDIR /app
RUN mkdir -p target && \
    GOOS=linux go build -o target/ender_chest main.go


FROM alpine:3.19 as prod

WORKDIR /app
COPY --from=0 /app/target/littlebox .

ENTRYPOINT ["/app/littlebox"]