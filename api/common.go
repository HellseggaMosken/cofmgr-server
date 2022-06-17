package api

import (
	"net/http"

	"cofmgr/logger"
	"cofmgr/model"

	"github.com/gin-gonic/gin"
)

// @Tags Test
// @Summary Ping
// @Description Check server status
// @Router /ping [get]
// @Produce json
// @Success 200 {} string "pong"
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

// 获取当前用户, 通过 AuthRequired 中间件后可调用此方法
func currentUser(c *gin.Context) *model.User {
	if value, exists := c.Get("user"); exists && value != nil {
		if u, ok := value.(*model.User); ok {
			return u
		}
		logger.Panic("in currentUser:", "can't convert")
	}
	logger.Panic("in currentUser:", "not found")
	return nil
}

// 获取当前用户, 通过AuthRequired中间件后可调用此方法
func currentAdmin(c *gin.Context) *model.Admin {
	if value, exists := c.Get("admin"); exists && value != nil {
		if a, ok := value.(*model.Admin); ok {
			return a
		}
		logger.Panic("in currentAdmin:", "can't convert")
	}
	logger.Panic("in currentAdmin:", "not found")
	return nil
}
