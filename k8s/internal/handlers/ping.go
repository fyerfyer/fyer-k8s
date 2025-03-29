package handlers

import (
	"k8s-demo/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type pingHandler struct {
	pingService services.PingService
}

func NewPingHandler() *pingHandler {
	return &pingHandler{
		pingService: services.NewPingService(),
	}
}

func(ph *pingHandler) Ping(c *gin.Context) {
	ph.pingService.Ping()

	c.JSON(http.StatusOK, gin.H{
		"message": "ping success",
	})
}