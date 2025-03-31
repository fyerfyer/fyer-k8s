package services

import (
    "context"
    "k8s-demo/internal/kube"
    "k8s-demo/internal/models"
    
    corev1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
)

type PodService interface {
    ListPods(ctx context.Context, namespace string) (*models.PodListResponse, error)
    GetPodDetail(ctx context.Context, namespace, name string) (*models.PodInfo, error)
}

type podService struct {
    kubeClientSet *kubernetes.Clientset
}

func NewPodService() PodService {
    return &podService{
        kubeClientSet: kube.GetKubeClientSet(),
    }
}

// ListPods 获取指定命名空间的Pod列表
func (ps *podService) ListPods(ctx context.Context, namespace string) (*models.PodListResponse, error) {
    var podList *corev1.PodList
    var err error

    if namespace == "" || namespace == "all" {
        podList, err = ps.kubeClientSet.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
    } else {
        podList, err = ps.kubeClientSet.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
    }

    if err != nil {
        return nil, err
    }

    result := &models.PodListResponse{
        Pods:  make([]models.PodInfo, 0, len(podList.Items)),
        Total: len(podList.Items),
    }

    for _, pod := range podList.Items {
        podInfo := convertPodToInfo(&pod)
        result.Pods = append(result.Pods, podInfo)
    }

    return result, nil
}

// GetPodDetail 获取单个Pod的详细信息
func (ps *podService) GetPodDetail(ctx context.Context, namespace, name string) (*models.PodInfo, error) {
    pod, err := ps.kubeClientSet.CoreV1().Pods(namespace).Get(ctx, name, metav1.GetOptions{})
    if err != nil {
        return nil, err
    }

    podInfo := convertPodToInfo(pod)
    return &podInfo, nil
}

// convertPodToInfo 将K8s Pod对象转换为自定义PodInfo对象
func convertPodToInfo(pod *corev1.Pod) models.PodInfo {
    podInfo := models.PodInfo{
        Name:         pod.Name,
        Namespace:    pod.Namespace,
        Status:       string(pod.Status.Phase),
        PodIP:        pod.Status.PodIP,
        NodeName:     pod.Spec.NodeName,
        CreationTime: pod.CreationTimestamp.Time,
        Labels:       pod.Labels,
        Containers:   make([]models.ContainerInfo, 0, len(pod.Spec.Containers)),
    }

    // 处理容器状态
    for i, container := range pod.Spec.Containers {
        containerStatus := models.ContainerInfo{
            Name:  container.Name,
            Image: container.Image,
        }

        // 尝试查找对应的容器状态
        if i < len(pod.Status.ContainerStatuses) {
            status := pod.Status.ContainerStatuses[i]
            containerStatus.Ready = status.Ready
            containerStatus.RestartCount = status.RestartCount

            // 确定容器状态
            if status.State.Running != nil {
                containerStatus.State = "Running"
            } else if status.State.Waiting != nil {
                containerStatus.State = "Waiting"
            } else if status.State.Terminated != nil {
                containerStatus.State = "Terminated"
            }
        }

        podInfo.Containers = append(podInfo.Containers, containerStatus)
    }

    return podInfo
}