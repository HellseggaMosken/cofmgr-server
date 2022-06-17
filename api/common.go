package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
