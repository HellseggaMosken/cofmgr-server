package middleware

import (
	"cofmgr/model"
	"cofmgr/service"
	"cofmgr/service/authservice"

	"github.com/gin-gonic/gin"
)

type AuthRequiredType int

const (
	AuthRequiredAdmin AuthRequiredType = iota
	AuthRequiredUser
	AuthRequiredAny // user or admin
)

func AuthRequired(requiredType AuthRequiredType) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		isAdmin, id, expired, valid := authservice.ParseToken(token)

		if !valid {
			c.AbortWithStatus(service.StatusNoAuth)
			return
		}
		if expired {
			c.AbortWithStatusJSON(service.StatusNoAuth, "token expired")
			return
		}

		if requiredType == AuthRequiredAdmin && !isAdmin {
			c.AbortWithStatusJSON(service.StatusNoAuth, "need admin role")
			return
		}

		if requiredType == AuthRequiredUser && isAdmin {
			c.AbortWithStatusJSON(service.StatusNoAuth, "need user role")
			return
		}

		if isAdmin {
			if admin := model.FindAdminWithID(id); admin == nil {
				c.AbortWithStatusJSON(service.StatusNoAuth, "no admin found")
			} else {
				c.Set("admin", admin)
			}
		} else {
			if user := model.FindUserWithID(id); user == nil {
				c.AbortWithStatusJSON(service.StatusNoAuth, "no user found")
			} else {
				c.Set("user", user)
			}
		}
		c.Next()
		return
	}
}
