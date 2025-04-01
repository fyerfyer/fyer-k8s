package handlers

import (
    "k8s-demo/internal/services"
    "k8s-demo/internal/models"
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

// DeletePod 删除指定的Pod
func (ph *podHandler) DeletePod(c *gin.Context) {
    namespace := c.Param("namespace")
    name := c.Param("name")

    if namespace == "" || name == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "namespace and pod name are required",
        })
        return
    }

    err := ph.podService.DeletePod(c.Request.Context(), namespace, name)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Pod deleted successfully",
        "namespace": namespace,
        "name": name,
    })
}

// CreatePod 创建新的Pod
func (ph *podHandler) CreatePod(c *gin.Context) {
    var podRequest models.PodCreateRequest

    // 绑定请求参数
    if err := c.ShouldBindJSON(&podRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request format: " + err.Error(),
        })
        return
    }

    // 验证必需的字段
    if podRequest.Name == "" || podRequest.Namespace == "" || len(podRequest.Containers) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Pod name, namespace and at least one container are required",
        })
        return
    }

    // 调用服务创建Pod
    pod, err := ph.podService.CreatePod(c.Request.Context(), &podRequest)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to create pod: " + err.Error(),
        })
        return
    }

    c.JSON(http.StatusCreated, pod)
}