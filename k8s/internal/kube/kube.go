package kube

import (
    "fmt"
    "k8s-demo/internal/config"
    "os"
    "path/filepath"
    "sync"

    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
    "k8s.io/client-go/tools/clientcmd"
)

var (
    kubeClientSet *kubernetes.Clientset
    once         sync.Once
)

// GetKubeClientSet 返回 Kubernetes 客户端
func GetKubeClientSet() *kubernetes.Clientset {
    once.Do(func() {
        // 初始化客户端
        clientset, err := createKubeClient()
        if err != nil {
            panic(fmt.Sprintf("Failed to create Kubernetes client: %v", err))
        }
        kubeClientSet = clientset
    })
    return kubeClientSet
}

// createKubeClient 尝试使用多种方法创建 Kubernetes 客户端
func createKubeClient() (*kubernetes.Clientset, error) {
    configMethods := []struct {
        name string
        fn   func() (*rest.Config, error)
    }{
        {"KUBECONFIG environment variable", tryEnvKubeConfig},
        {"Path specified in configuration file", tryConfigFileKubeConfig},
        {"Default path ~/.kube/config", tryDefaultKubeConfig},
        {"In-cluster configuration", tryInClusterConfig},
    }

    var lastErr error
    for _, method := range configMethods {
        fmt.Printf("Attempting to load Kubernetes configuration using %s...\n", method.name)
        config, err := method.fn()
        if err != nil {
            fmt.Printf("  - Failed: %v\n", err)
            lastErr = err
            continue
        }

        // 成功获取配置，创建客户端
        fmt.Printf("  - Success!\n")
        return kubernetes.NewForConfig(config)
    }

    return nil, fmt.Errorf("failed to load a valid Kubernetes configuration: %v", lastErr)
}

// tryEnvKubeConfig 尝试从环境变量 KUBECONFIG 中加载配置
func tryEnvKubeConfig() (*rest.Config, error) {
    kubeconfigPath := os.Getenv("KUBECONFIG")
    if kubeconfigPath == "" {
        return nil, fmt.Errorf("KUBECONFIG environment variable is not set")
    }
    
    return loadKubeConfig(kubeconfigPath)
}

// tryConfigFileKubeConfig 尝试从配置文件中加载配置
func tryConfigFileKubeConfig() (*rest.Config, error) {
    kubeconfigPath := config.GlobalConfig.Kubernetes.ConfigPath
    if kubeconfigPath == "" {
        return nil, fmt.Errorf("kubernetes.configPath is not specified in the configuration file")
    }
    
    // 如果路径以 ~ 开头，则将其展开为用户的主目录
    if len(kubeconfigPath) > 0 && kubeconfigPath[0] == '~' {
        home, err := os.UserHomeDir()
        if err != nil {
            return nil, fmt.Errorf("failed to get user's home directory: %v", err)
        }
        kubeconfigPath = filepath.Join(home, kubeconfigPath[1:])
    }
    
    return loadKubeConfig(kubeconfigPath)
}

// tryDefaultKubeConfig 尝试从默认路径 ~/.kube/config 中加载配置
func tryDefaultKubeConfig() (*rest.Config, error) {
    home, err := os.UserHomeDir()
    if err != nil {
        return nil, fmt.Errorf("failed to get user's home directory: %v", err)
    }
    
    kubeconfigPath := filepath.Join(home, ".kube", "config")
    return loadKubeConfig(kubeconfigPath)
}

// tryInClusterConfig 尝试使用集群内配置
func tryInClusterConfig() (*rest.Config, error) {
    return rest.InClusterConfig()
}

// loadKubeConfig 加载指定路径的 kubeconfig 文件
func loadKubeConfig(path string) (*rest.Config, error) {
    _, err := os.Stat(path)
    if err != nil {
        return nil, fmt.Errorf("kubeconfig file does not exist or is not accessible: %v", err)
    }
    
    fmt.Printf("  - Attempting to load configuration from %s\n", path)
    return clientcmd.BuildConfigFromFlags("", path)
}