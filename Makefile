GOCMD=/usr/local/bin/go

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPROXY=https://goproxy.io $(GOCMD) build -o ./drone-aliyun-oss/drone-aliyun-oss ./drone-aliyun-oss/main.go
