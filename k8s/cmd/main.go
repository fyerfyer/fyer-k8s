package main

import (
    "fmt"
    "k8s-demo/internal/config"
    "k8s-demo/internal/router"
    "log"
)

func main() {
    // 加载配置文件
    if err := config.LoadConfig("./internal/config/config.yaml"); err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }

    r := router.SetUpRouter()
    serverAddr := fmt.Sprintf("%s:%d", 
        config.GlobalConfig.Server.Host, 
        config.GlobalConfig.Server.Port)

    if err := r.Run(serverAddr); err != nil {
        log.Fatalf("Starting serving error: %v", err)
    }
}