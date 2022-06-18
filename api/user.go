package api

import (
	"cofmgr/service/userservice"

	"github.com/gin-gonic/gin"
)

// @Tags User
// @Summary User Register
// @Description Register a new user
// @Router /users [post]
// @Param body body userservice.RegisterService true "HTTP Body"
// @Produce json
// @Success 200 {object} serializer.User "Sucess"
// @Failure 406 {string} string          "Email exists"
func UserRegister(c *gin.Context) {
	var s userservice.RegisterService
	if !bind(c, &s) {
		return
	}
	user, code := s.Register()
	c.JSON(code, user)
}

// @Tags User
// @Summary User Update
// @Description Update user info
// @Router /users [put]
// @Param token body string                    true "User Token"
// @Param body  body userservice.UpdateService true "HTTP Body"
// @Produce json
// @Success 200 {object} serializer.User "Sucess"
// @Failure 406 {string} string          "Email exists"
func UserUpdate(c *gin.Context) {
	var s userservice.UpdateService
	if !bind(c, &s) {
		return
	}
	user, code := s.Update(currentUser(c))
	c.JSON(code, user)
}

// @Tags User
// @Summary User Show
// @Description Show user info
// @Router /users/{user_id} [get]
// @Param token body string true "Admin Token"
// @Produce json
// @Success 200 {object} serializer.User "Sucess"
// @Failure 404 {string} string          "Not found"
func UserShow(c *gin.Context) {
	user, code := userservice.Show(c.Param("id"))
	c.JSON(code, user)
}

// @Tags User
// @Summary User List
// @Description List all users
// @Router /users [get]
// @Param token body string true "Admin Token"
// @Produce json
// @Success 200 {array}  serializer.User "Sucess"
func UserList(c *gin.Context) {
	users, code := userservice.List()
	c.JSON(code, users)
}
