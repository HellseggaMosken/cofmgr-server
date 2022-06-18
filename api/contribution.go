package api

import (
	"cofmgr/service/ctbservice"

	"github.com/gin-gonic/gin"
)

func ContributionCreate(c *gin.Context) {
	var s ctbservice.CreateService
	if !bind(c, &s) {
		return
	}
	ctb, code := s.Create(currentUser(c))
	c.JSON(code, ctb)
}

func ContributionUpdate(c *gin.Context) {
	var s ctbservice.UpdateService
	if !bind(c, &s) {
		return
	}
	ctb, code := s.Update()
	c.JSON(code, ctb)
}

func ContributionShow(c *gin.Context) {
	ctb, code := ctbservice.Show(c.Param("id"))
	c.JSON(code, ctb)
}

func ContributionListWithUser(c *gin.Context) {
	cofs, code := ctbservice.ListWithUser(currentUser(c))
	c.JSON(code, cofs)
}

func ContributionListWithConference(c *gin.Context) {
	cofs, code := ctbservice.ListWithConference(c.Param("cof_id"))
	c.JSON(code, cofs)
}
