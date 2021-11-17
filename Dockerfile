FROM golang:1.16.5-alpine3.13 as builder

WORKDIR /go/src
ADD . .

RUN GO111MODULE=on GOPROXY="https://goproxy.cn" CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o /go/bin/drone-aliyun-oss .

FROM alpine:3.11
RUN echo "http://mirrors.aliyun.com/alpine/v3.11/main" > /etc/apk/repositories
RUN echo "http://mirrors.aliyun.com/alpine/v3.11/community" >> /etc/apk/repositories
RUN apk add -U tzdata \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime \
    && apk add ca-certificates \
    && rm -rf /var/cache/apk/*

WORKDIR /go/bin

COPY --from=builder /go/bin/drone-aliyun-oss .
ENTRYPOINT ["/go/bin/drone-aliyun-oss"]