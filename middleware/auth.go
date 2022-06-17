package middleware

import (
	"net/http"

	"cofmgr/model"

	"github.com/gin-gonic/gin"
)

func AuthRequired(isAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = c.GetHeader("token")
		notAuthed := false
		expired := false
		if notAuthed {
			c.AbortWithStatus(http.StatusForbidden)
		}
		if expired {
			c.AbortWithStatusJSON(http.StatusForbidden, "expired")
		}
		isAdmin := true
		if isAdmin {
			c.Set("admin", model.Admin{})
		} else {
			c.Set("user", model.User{})
		}
		c.Next()
		return
	}
}
