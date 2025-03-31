package router

import (
    "k8s-demo/internal/handlers"

    "github.com/gin-gonic/gin"
)

func setupPodAPIRoutes(r *gin.Engine) {
    podHandler := handlers.NewPodHandler()
    
    // Pod API 路由组
    podGroup := r.Group("/api/v1/pods")
    {
        // 获取Pod列表
        podGroup.GET("", podHandler.ListPods)
        
        // 获取Pod详情
        podGroup.GET("/:namespace/:name", podHandler.GetPodDetail)
    }
}