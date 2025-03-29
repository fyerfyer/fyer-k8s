package router

import (
	"k8s-demo/internal/handlers"

	"github.com/gin-gonic/gin"
)

func setupPingAPIRoutes(r *gin.Engine) {
	pingHandler := handlers.NewPingHandler()
	r.GET("/ping", pingHandler.Ping)
}