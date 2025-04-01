package models

import (
    "time"
)

// PodInfo 定义了Pod的基本信息
type PodInfo struct {
    Name           string            `json:"name"`
    Namespace      string            `json:"namespace"`
    Status         string            `json:"status"`
    PodIP          string            `json:"podIP"`
    NodeName       string            `json:"nodeName"`
    CreationTime   time.Time         `json:"creationTime"`
    Labels         map[string]string `json:"labels,omitempty"`
    Containers     []ContainerInfo   `json:"containers,omitempty"`
}

// ContainerInfo 定义了容器的基本信息
type ContainerInfo struct {
    Name          string `json:"name"`
    Image         string `json:"image"`
    Ready         bool   `json:"ready"`
    RestartCount  int32  `json:"restartCount"`
    State         string `json:"state"`
}

// PodListResponse 定义了Pod列表的响应格式
type PodListResponse struct {
    Pods  []PodInfo `json:"pods"`
    Total int       `json:"total"`
}

// PodCreateRequest 定义了创建Pod的请求结构
type PodCreateRequest struct {
    Name       string                 `json:"name" binding:"required"`
    Namespace  string                 `json:"namespace" binding:"required"`
    Labels     map[string]string      `json:"labels,omitempty"`
    Containers []ContainerCreateSpec  `json:"containers" binding:"required,min=1"`
}

// ContainerCreateSpec 定义了创建容器的规格
type ContainerCreateSpec struct {
    Name     string             `json:"name" binding:"required"`
    Image    string             `json:"image" binding:"required"`
    Ports    []ContainerPort    `json:"ports,omitempty"`
    Env      map[string]string  `json:"env,omitempty"`
}

// ContainerPort 定义了容器端口配置
type ContainerPort struct {
    Name          string `json:"name,omitempty"`
    ContainerPort int32  `json:"containerPort" binding:"required"`
    Protocol      string `json:"protocol,omitempty"` // 默认为TCP，可选值：TCP, UDP, SCTP
}