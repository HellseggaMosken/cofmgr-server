package api

import (
	"cofmgr/service/cofservice"

	"github.com/gin-gonic/gin"
)

// @Tags Conference
// @Summary Create Conference
// @Description Create a conference
// @Router /conferences [post]
// @Param token body string                   true "Admin Token"
// @Param body  body cofservice.CreateService true "HTTP Body"
// @Produce json
// @Success 200 {object} serializer.Conference "Sucess"
// @Failure 403 {string} string                "No token or token invalid or expired"
func ConferenceCreate(c *gin.Context) {
	var s cofservice.CreateService
	if !bind(c, &s) {
		return
	}
	cof, code := s.Create()
	c.JSON(code, cof)
}

// @Tags Conference
// @Summary Update Conference
// @Description Update a conference
// @Router /conferences [put]
// @Param token body string                   true "Admin Token"
// @Param body  body cofservice.UpdateService true "HTTP Body"
// @Produce json
// @Success 200 {object} serializer.Conference "Sucess"
// @Failure 403 {string} string                "No token or token invalid or expired"
func ConferenceUpdate(c *gin.Context) {
	var s cofservice.UpdateService
	if !bind(c, &s) {
		return
	}
	cof, code := s.Update()
	c.JSON(code, cof)
}

// @Tags Conference
// @Summary Show Conference
// @Description Show a conference info
// @Router /conferences/{conference_id} [get]
// @Produce json
// @Success 200 {object} serializer.Conference "Sucess"
// @Failure 404 {string} string                "Not found"
func ConferenceShow(c *gin.Context) {
	cof, code := cofservice.Show(c.Param("id"))
	c.JSON(code, cof)
}

// @Tags Conference
// @Summary List All Conferences
// @Description List all conferences
// @Router /conferences [get]
// @Produce json
// @Success 200 {array}  serializer.Conference "Sucess"
func ConferenceList(c *gin.Context) {
	cofs, code := cofservice.List()
	c.JSON(code, cofs)
}
