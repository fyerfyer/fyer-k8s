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

# 在builder阶段创建Kubernetes配置目录
RUN mkdir -p /root/.kube

FROM gcr.io/distroless/static-debian11

# 复制二进制文件和必要的配置
COPY --from=builder /app/k8s-app /app/
COPY --from=builder /app/internal/config/config.yaml /app/internal/config/
# 复制预先创建的目录
COPY --from=builder /root/.kube /root/.kube

# 工作目录
WORKDIR /app

# 暴露服务端口
EXPOSE 8081

# 运行应用
CMD ["/app/k8s-app"]