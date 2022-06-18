package api

import (
	"cofmgr/service/userservice"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var s userservice.RegisterService
	if !bind(c, &s) {
		return
	}
	user, code := s.Register()
	c.JSON(code, user)
}

func UserUpdate(c *gin.Context) {
	var s userservice.UpdateService
	if !bind(c, &s) {
		return
	}
	user, code := s.Update(currentUser(c))
	c.JSON(code, user)
}

func UserShow(c *gin.Context) {
	user, code := userservice.Show(c.Param("id"))
	c.JSON(code, user)
}

func UserList(c *gin.Context) {
	users, code := userservice.List()
	c.JSON(code, users)
}
