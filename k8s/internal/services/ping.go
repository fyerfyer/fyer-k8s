package services

import "fmt"

type PingService interface {
	Ping () error 
}

type pingService struct {}

func NewPingService() *pingService {
	return &pingService{}
}

func (ps *pingService) Ping() error {
	fmt.Println("Pong")
	return nil
}