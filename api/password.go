package api

import (
	"cofmgr/service/passwordservice"

	"github.com/gin-gonic/gin"
)

func PasswordChangeForUser(c *gin.Context) {
	var s passwordservice.PasswordService
	if !bind(c, &s) {
		return
	}
	code := s.ChangeUserPassword(currentUser(c))
	c.Status(code)
}

func PasswordChangeForAdmin(c *gin.Context) {
	var s passwordservice.PasswordService
	if !bind(c, &s) {
		return
	}
	code := s.ChangeAdminPassword(currentAdmin(c))
	c.Status(code)
}
