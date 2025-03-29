package kube

import (
	"k8s-demo/internal/config"
	"sync"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
    kubeClientSet *kubernetes.Clientset
    once         sync.Once
)

func GetKubeClientSet() *kubernetes.Clientset {
    once.Do(func() {
        var err error
        kubeconfig := config.GlobalConfig.Kubernetes.ConfigPath
        config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
        if err != nil {
            panic(err.Error())
        }

        kubeClientSet, err = kubernetes.NewForConfig(config)
        if err != nil {
            panic(err.Error())
        }
    })
    return kubeClientSet
}