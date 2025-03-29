package services

import (
	"context"
	"log"
	"k8s-demo/internal/kube"
	"k8s.io/client-go/kubernetes"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodService interface {
	ListPod(ctx context.Context) error 
}

type podService struct {
	kubeClientSet *kubernetes.Clientset
}

func NewPodService () *podService {
	return &podService{
		kubeClientSet: kube.GetKubeClientSet(),
	}
}

func (ps *podService) ListPod(ctx context.Context) error {
	list, err := ps.kubeClientSet.CoreV1().
		Namespaces().List(ctx, metav1.ListOptions{})
	
	if err != nil {
		return err 
	}

	for _, item := range list.Items {
		log.Printf("item namespace: %s, name: %v", item.Namespace, item.Name)
	}

	return nil
}