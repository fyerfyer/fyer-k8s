package config

import (
    "fmt"
    "os"

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
    data, err := os.ReadFile(configPath)
    if err != nil {
        return fmt.Errorf("error reading config file: %v", err)
    }

    err = yaml.Unmarshal(data, &GlobalConfig)
    if err != nil {
        return fmt.Errorf("error parsing config file: %v", err)
    }

    return nil
}