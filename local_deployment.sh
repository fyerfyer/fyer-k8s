#!/bin/bash

# 设置变量
GITHUB_USER="fyerfyer"  # docker-compose.yml中使用的用户名
VERSION="latest"  # 默认使用最新版本

# 检查参数
if [ "$1" == "down" ]; then
    echo "停止并移除容器..."
    docker-compose down
    exit 0
fi

if [ "$1" != "" ]; then
    VERSION="$1"
    echo "使用指定版本: $VERSION"
else
    echo "使用默认版本: $VERSION"
fi

# 登录到GitHub容器注册表
echo "登录到GitHub容器注册表..."
echo "请输入你的GitHub个人访问令牌:"
read -s GITHUB_TOKEN
echo $GITHUB_TOKEN | docker login ghcr.io -u $GITHUB_USER --password-stdin

# 创建临时docker-compose文件并替换版本号
echo "创建临时部署文件..."
cp docker-compose.yml docker-compose.local.yml

# 如果版本不是latest，替换版本号
if [ "$VERSION" != "latest" ]; then
    sed -i "s|:latest|:${VERSION}|g" docker-compose.local.yml
    echo "版本已更新为: $VERSION"
fi

# 拉取和启动容器
echo "拉取最新镜像..."
docker-compose -f docker-compose.local.yml pull

echo "启动容器..."
docker-compose -f docker-compose.local.yml up -d

echo "部署完成！应用正在运行:"
echo "- 前端: http://localhost:8080"
echo "- 后端: http://localhost:8081"
echo "容器已启动！使用 './local_deployment.sh down' 来停止它们"