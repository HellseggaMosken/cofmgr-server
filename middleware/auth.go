package middleware

import (
	"net/http"

	"cofmgr/model"
	"cofmgr/service/authservice"

	"github.com/gin-gonic/gin"
)

func AuthRequired(isAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		isAdmin, id, expired, valid := authservice.ParseToken(token)

		if !valid {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		if expired {
			c.AbortWithStatusJSON(http.StatusForbidden, "token expired")
			return
		}

		if isAdmin {
			if admin := model.FindAdminWithID(id); admin == nil {
				c.AbortWithStatusJSON(http.StatusForbidden, "no admin found")
			} else {
				c.Set("admin", admin)
			}
		} else {
			if user := model.FindAdminWithID(id); user == nil {
				c.AbortWithStatusJSON(http.StatusForbidden, "no user found")
			} else {
				c.Set("admin", user)
			}
		}
		c.Next()
		return
	}
}
