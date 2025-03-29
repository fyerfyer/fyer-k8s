package router

import (
	"k8s-demo/internal/handlers"

	"github.com/gin-gonic/gin"
)

func setupPodAPIRoutes(r *gin.Engine) {
	podHandler := handlers.NewPodHandler()
	r.GET("/pod/list", podHandler.ListPod)
}