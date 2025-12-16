package handler

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tandyys/fds-service/internal/dto"
)

func PingGetHandler(c *gin.Context) {
	c.JSON(http.StatusOK, dto.PingResponse{
		Status:  http.StatusOK,
		Message: "What's up! How's life my G?",
	})
}

func PingGetRandomNumberHandler(c *gin.Context) {
	randomNumber := rand.Int()

	c.JSON(http.StatusOK, dto.PingRandomNumberResponse{
		Status:       http.StatusOK,
		RandomNumber: randomNumber,
	})
}
