package router

import "github.com/gin-gonic/gin"

func SetUpRoutes() *gin.Engine {
	r := gin.Default()

	registerPingRoutes(r)

	return r
}
