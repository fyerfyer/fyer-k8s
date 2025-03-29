FROM node:18-alpine AS build-stage

# 设置工作目录
WORKDIR /app

# 复制配置文件
COPY package*.json ./

# 下载依赖
RUN npm ci

COPY . .

# 构建应用
RUN npm run build

# 使用 nginx 作为生产环境的基础镜像
FROM nginx:stable-alpine AS production-stage

# 复制构建好的文件到 nginx 的 html 目录
COPY --from=build-stage /app/dist /usr/share/nginx/html

# 复制 nginx 配置文件
COPY ./nginx.conf /etc/nginx/conf.d/default.conf

# 创建一个非 root 用户来运行容器
RUN addgroup -g 1000 -S appgroup && \
    adduser -u 1000 -S appuser -G appgroup

# 设置权限
RUN chown -R appuser:appgroup /usr/share/nginx/html && \
    chown -R appuser:appgroup /var/cache/nginx && \
    chown -R appuser:appgroup /var/log/nginx && \
    touch /var/run/nginx.pid && \
    chown -R appuser:appgroup /var/run/nginx.pid

# 切换到非 root 用户
USER appuser

# Nginx 默认使用 80 端口
EXPOSE 80

# 使用 wget 来检查 Nginx 是否正常运行
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget -q --spider http://localhost:80/ || exit 1

# 启动 Nginx
CMD ["nginx", "-g", "daemon off;"]