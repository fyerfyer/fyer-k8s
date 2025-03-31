package config

import (
    "fmt"
    "os"
    "path/filepath"

    "gopkg.in/yaml.v3"
)

type Config struct {
    Server struct {
        Port int    `yaml:"port"`
        Host string `yaml:"host"`
    } `yaml:"server"`
    Kubernetes struct {
        ConfigPath string `yaml:"configPath"`
    } `yaml:"kubernetes"`
}

var GlobalConfig Config

func LoadConfig(configPath string) error {
    // 设置默认配置
    setDefaultConfig()

    data, err := os.ReadFile(configPath)
    if err != nil {
        // 如果配置文件不存在，使用默认配置
        if os.IsNotExist(err) {
            fmt.Printf("Warning: Config file %s not found, using default configuration\n", configPath)
            return nil
        }
        return fmt.Errorf("error reading config file: %v", err)
    }

    err = yaml.Unmarshal(data, &GlobalConfig)
    if err != nil {
        return fmt.Errorf("error parsing config file: %v", err)
    }

    // 从环境变量覆盖配置
    loadFromEnv()

    return nil
}

// 设置默认配置
func setDefaultConfig() {
    GlobalConfig.Server.Port = 8081
    GlobalConfig.Server.Host = "0.0.0.0"
    
    // 默认使用$HOME/.kube/config
    home, err := os.UserHomeDir()
    if err == nil {
        GlobalConfig.Kubernetes.ConfigPath = filepath.Join(home, ".kube", "config")
    }
}

// 从环境变量加载配置
func loadFromEnv() {
    // 服务器端口
    if port := os.Getenv("SERVER_PORT"); port != "" {
        fmt.Sscanf(port, "%d", &GlobalConfig.Server.Port)
    }
    
    // 服务器主机
    if host := os.Getenv("SERVER_HOST"); host != "" {
        GlobalConfig.Server.Host = host
    }
    
    // Kubernetes配置
    if kubeconfig := os.Getenv("KUBECONFIG"); kubeconfig != "" {
        GlobalConfig.Kubernetes.ConfigPath = kubeconfig
    }
}