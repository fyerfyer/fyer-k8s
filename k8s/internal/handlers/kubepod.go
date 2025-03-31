package handlers

import (
    "k8s-demo/internal/services"
    "net/http"

    "github.com/gin-gonic/gin"
)

type podHandler struct {
    podService services.PodService
}

func NewPodHandler() *podHandler {
    return &podHandler{
        podService: services.NewPodService(),
    }
}

// ListPods 获取Pod列表
func (ph *podHandler) ListPods(c *gin.Context) {
    namespace := c.Query("namespace")

    pods, err := ph.podService.ListPods(c.Request.Context(), namespace)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, pods)
}

// GetPodDetail 获取单个Pod详情
func (ph *podHandler) GetPodDetail(c *gin.Context) {
    namespace := c.Param("namespace")
    name := c.Param("name")

    if namespace == "" || name == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "namespace and pod name are required",
        })
        return
    }

    pod, err := ph.podService.GetPodDetail(c.Request.Context(), namespace, name)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, pod)
}