package api

import (
	"cofmgr/service/ctbservice"

	"github.com/gin-gonic/gin"
)

// @Tags Contribution
// @Summary Create Contribution
// @Description Create a contribution
// @Router /contributions [post]
// @Param token body string                   true "User Token"
// @Param body  body ctbservice.CreateService true "HTTP Body"
// @Produce json
// @Success 200 {object} serializer.Contribution "Sucess"
// @Failure 403 {string} string                  "No token or token invalid or expired"
func ContributionCreate(c *gin.Context) {
	var s ctbservice.CreateService
	if !bind(c, &s) {
		return
	}
	ctb, code := s.Create(currentUser(c))
	c.JSON(code, ctb)
}

// @Tags Contribution
// @Summary Update Contribution
// @Description Update a contribution
// @Router /contributions [put]
// @Param token body string                   true "User Token"
// @Param body  body ctbservice.UpdateService true "HTTP Body"
// @Produce json
// @Success 200 {object} serializer.Contribution "Sucess"
// @Failure 403 {string} string                  "No token or token invalid or expired"
func ContributionUpdate(c *gin.Context) {
	var s ctbservice.UpdateService
	if !bind(c, &s) {
		return
	}
	ctb, code := s.Update()
	c.JSON(code, ctb)
}

// @Tags Contribution
// @Summary Show Contribution
// @Description Show a contribution info
// @Router /contributions/common/{contribution_id} [get]
// @Produce json
// @Success 200 {object} serializer.Contribution "Sucess"
// @Failure 404 {string} string                  "Not found"
func ContributionShow(c *gin.Context) {
	ctb, code := ctbservice.Show(c.Param("id"))
	c.JSON(code, ctb)
}

// @Tags Contribution
// @Summary List All Contributions of Current Contributor
// @Description List all contributions of the current contributor(user)
// @Router /contributions/user [get]
// @Param token body string true "User Token"
// @Produce json
// @Success 200 {array} serializer.Contribution "Sucess"
// @Failure 403 {string} string                  "No token or token invalid or expired"
func ContributionListWithUser(c *gin.Context) {
	cofs, code := ctbservice.ListWithUser(currentUser(c))
	c.JSON(code, cofs)
}

// @Tags Contribution
// @Summary List All Contributions of a Conference
// @Description List all contributions of a conference
// @Router /contributions/conference/{conference_id} [get]
// @Param token body string true "User Token"
// @Produce json
// @Success 200 {array} serializer.Contribution "Sucess"
// @Failure 403 {string} string                 "No token or token invalid or expired"
func ContributionListWithConference(c *gin.Context) {
	cofs, code := ctbservice.ListWithConference(c.Param("cof_id"))
	c.JSON(code, cofs)
}
