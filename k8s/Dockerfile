# 使用较小的golang镜像
FROM golang:1.23-alpine AS builder

WORKDIR /app

# 安装必要的构建工具
RUN apk add --no-cache git

# 复制go模块文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建可执行文件，添加一些优化选项来减小二进制文件大小
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /app/k8s-app ./cmd

# 使用scratch作为最终基础镜像，这是一个空镜像
FROM scratch

# 复制证书，确保HTTPS请求能够工作
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# 创建工作目录
WORKDIR /app

# 仅复制编译好的二进制文件和必要的配置
COPY --from=builder /app/k8s-app .
COPY --from=builder /app/internal/config/config.yaml ./internal/config/

# 创建Kubernetes配置目录
WORKDIR /app/.kube

# 复制 kubernetes 配置文件
COPY --from=builder /app/.kube .

WORKDIR /app

# 暴露服务端口
EXPOSE 8081

# 运行应用
CMD ["/app/k8s-app"]