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
		isAdmin, _, expired, valid := authservice.ParseToken(token)

		if !valid {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		if expired {
			c.AbortWithStatusJSON(http.StatusForbidden, "token expired")
			return
		}

		if isAdmin {
			c.Set("admin", model.Admin{})
		} else {
			c.Set("user", model.User{})
		}
		c.Next()
		return
	}
}
