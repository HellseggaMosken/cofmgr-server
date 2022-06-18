package middleware

import (
	"cofmgr/model"
	"cofmgr/service"
	"cofmgr/service/authservice"

	"github.com/gin-gonic/gin"
)

func AuthRequired(isAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		_isAdmin, id, expired, valid := authservice.ParseToken(token)

		if !valid {
			c.AbortWithStatus(service.StatusNoAuth)
			return
		}
		if expired {
			c.AbortWithStatusJSON(service.StatusNoAuth, "token expired")
			return
		}

		if isAdmin != _isAdmin {
			c.AbortWithStatusJSON(service.StatusNoAuth, "wrong role")
			return
		}

		if isAdmin {
			if admin := model.FindAdminWithID(id); admin == nil {
				c.AbortWithStatusJSON(service.StatusNoAuth, "no admin found")
			} else {
				c.Set("admin", admin)
			}
		} else {
			if user := model.FindAdminWithID(id); user == nil {
				c.AbortWithStatusJSON(service.StatusNoAuth, "no user found")
			} else {
				c.Set("admin", user)
			}
		}
		c.Next()
		return
	}
}
