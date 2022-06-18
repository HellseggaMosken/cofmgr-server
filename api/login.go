package api

import (
	"cofmgr/service/loginservice"

	"github.com/gin-gonic/gin"
)

// @Tags Login
// @Summary User Login
// @Description User login
// @Router /login/user [post]
// @Param body  body loginservice.LoginService true "HTTP Body"
// @Produce json
// @Success 200 {object} serializer.User "Actually, is {"user": User, "token": string}"
// @Failure 401 {string} string          "Wrong password or account"
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

// @Tags Login
// @Summary Admin Login
// @Description Admin login
// @Router /login/admin [post]
// @Param body  body loginservice.LoginService true "HTTP Body"
// @Produce json
// @Success 200 {object} serializer.User "Actually, is {"admin": User, "token": string}"
// @Failure 401 {string} string          "Wrong password or account"
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
