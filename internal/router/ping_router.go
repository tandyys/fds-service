package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tandyys/fds-service/internal/handler"
)

func registerPingRoutes(r *gin.Engine) {
	r.GET("/ping", handler.PingGetHandler)
	r.GET("/random-number", handler.PingGetRandomNumberHandler)
}
