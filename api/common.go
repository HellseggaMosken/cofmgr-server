package api

import (
	"net/http"

	"cofmgr/logger"
	"cofmgr/model"

	"github.com/gin-gonic/gin"
)

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

// 获取当前用户, 通过 AuthRequired 中间件后可调用此方法
func currentUser(c *gin.Context) *model.User {
	if value, exists := c.Get("user"); exists && value != nil {
		if u, ok := value.(*model.User); ok {
			return u
		} else {
			logger.Warning("can't convert gin context value to *model.User when call currentUser")
		}
	} else {
		logger.Warning("not found user in gin context when call currentUser")
	}
	return nil
}

// 获取当前用户, 通过AuthRequired中间件后可调用此方法
func currentAdmin(c *gin.Context) *model.Admin {
	if value, exists := c.Get("admin"); exists && value != nil {
		if a, ok := value.(*model.Admin); ok {
			return a
		} else {
			logger.Warning("can't convert gin context value to *model.User when call currentUser")
		}
	} else {
		logger.Warning("not found user in gin context when call currentUser")
	}
	return nil
}
