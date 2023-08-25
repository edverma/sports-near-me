package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *server) pingHandler(c *gin.Context) {
	message := Ping()
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func Ping() string {
	return "pong"
}
