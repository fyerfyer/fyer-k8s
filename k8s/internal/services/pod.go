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
    DeletePod(ctx context.Context, namespace, name string) error
    CreatePod(ctx context.Context, podRequest *models.PodCreateRequest) (*models.PodInfo, error)
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

// DeletePod 删除指定的Pod
func (ps *podService) DeletePod(ctx context.Context, namespace, name string) error {
    deleteOptions := metav1.DeleteOptions{}
    return ps.kubeClientSet.CoreV1().Pods(namespace).Delete(ctx, name, deleteOptions)
}

// CreatePod 创建一个新的Pod
func (ps *podService) CreatePod(ctx context.Context, podRequest *models.PodCreateRequest) (*models.PodInfo, error) {
    // 创建容器列表
    containers := make([]corev1.Container, 0, len(podRequest.Containers))
    for _, c := range podRequest.Containers {
        container := corev1.Container{
            Name:  c.Name,
            Image: c.Image,
        }
        
        // 处理容器端口
        if len(c.Ports) > 0 {
            containerPorts := make([]corev1.ContainerPort, 0, len(c.Ports))
            for _, port := range c.Ports {
                containerPorts = append(containerPorts, corev1.ContainerPort{
                    Name:          port.Name,
                    ContainerPort: port.ContainerPort,
                    Protocol:      corev1.Protocol(port.Protocol),
                })
            }
            container.Ports = containerPorts
        }
        
        // 处理环境变量
        if len(c.Env) > 0 {
            envVars := make([]corev1.EnvVar, 0, len(c.Env))
            for k, v := range c.Env {
                envVars = append(envVars, corev1.EnvVar{
                    Name:  k,
                    Value: v,
                })
            }
            container.Env = envVars
        }
        
        containers = append(containers, container)
    }
    
    // 创建Pod对象
    pod := &corev1.Pod{
        ObjectMeta: metav1.ObjectMeta{
            Name:      podRequest.Name,
            Namespace: podRequest.Namespace,
            Labels:    podRequest.Labels,
        },
        Spec: corev1.PodSpec{
            Containers: containers,
        },
    }
    
    // 调用Kubernetes API创建Pod
    createdPod, err := ps.kubeClientSet.CoreV1().Pods(podRequest.Namespace).Create(ctx, pod, metav1.CreateOptions{})
    if err != nil {
        return nil, err
    }
    
    // 将创建的Pod转换为响应对象
    podInfo := convertPodToInfo(createdPod)
    return &podInfo, nil
}