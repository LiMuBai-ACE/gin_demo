# 基础镜像
FROM golang:1.13-alpine3.10 AS builder



# 工作区
WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download


# 复制仓库源文件到容器里
COPY . .

RUN go build -o main /app/main.go

FROM alpine:latest

COPY --from=builder /app/main .


# 容器环境变量设置，会覆盖默认的变量值
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE="on"
ENV ENV="production"

EXPOSE 4000

# 这里跟编译完的文件名一致
ENTRYPOINT  ["./main"]