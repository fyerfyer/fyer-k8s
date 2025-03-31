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