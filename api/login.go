package api

import (
	"cofmgr/service/loginservice"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	var s loginservice.LoginService
	if !bind(c, &s) {
		return
	}
	user, token, code := s.LoginUser()
	c.JSON(code, gin.H{
		"user":  user,
		"token": token,
	})
}

func LoginAdmin(c *gin.Context) {
	var s loginservice.LoginService
	if !bind(c, &s) {
		return
	}
	user, token, code := s.LoginAdmin()
	c.JSON(code, gin.H{
		"admin": user,
		"token": token,
	})
}
