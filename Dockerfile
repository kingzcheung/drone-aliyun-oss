FROM alpine:latest

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf /var/cache/apk/*

ADD drone-aliyun-oss/drone-aliyun-oss /bin/
ENTRYPOINT ["/bin/drone-aliyun-oss"]