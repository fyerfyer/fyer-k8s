package router

import "github.com/gin-gonic/gin"

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	setupPingAPIRoutes(r)
	setupPodAPIRoutes(r)
	return r
}