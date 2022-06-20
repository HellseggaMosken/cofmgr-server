package api

import (
	"cofmgr/service/passwordservice"

	"github.com/gin-gonic/gin"
)

// @Tags Password
// @Summary Change User Password
// @Description Change user password
// @Router /password/user [put]
// @Param token header string                          true "User Token"
// @Param body  body passwordservice.PasswordService true "HTTP Body"
// @Produce json
// @Success 200 {string} string "Sucess"
// @Failure 401 {string} string "Wrong password or account"
// @Failure 403 {string} string "No token or token invalid or expired"
func PasswordChangeForUser(c *gin.Context) {
	var s passwordservice.PasswordService
	if !bind(c, &s) {
		return
	}
	code := s.ChangeUserPassword(currentUser(c))
	c.Status(code)
}

// @Tags Password
// @Summary Change Admin Password
// @Description Change admin password
// @Router /password/admin [put]
// @Param token header string                          true "Admin Token"
// @Param body  body passwordservice.PasswordService true "HTTP Body"
// @Produce json
// @Success 200 {string} string "Sucess"
// @Failure 401 {string} string "Wrong password or account"
// @Failure 403 {string} string "No token or token invalid or expired"
func PasswordChangeForAdmin(c *gin.Context) {
	var s passwordservice.PasswordService
	if !bind(c, &s) {
		return
	}
	code := s.ChangeAdminPassword(currentAdmin(c))
	c.Status(code)
}
