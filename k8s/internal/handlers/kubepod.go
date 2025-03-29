package handlers

import (
	"k8s-demo/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type podHandler struct {
	posService services.PodService
}

func NewPodHandler() *podHandler {
	return &podHandler{
		posService: services.NewPodService(),
	}
}

func (ph *podHandler) ListPod(c *gin.Context) {
	if err := ph.posService.ListPod(c.Request.Context()); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H {
		"message": "pod list success",
	})
}