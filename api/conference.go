package api

import (
	"cofmgr/service/cofservice"

	"github.com/gin-gonic/gin"
)

func ConferenceCreate(c *gin.Context) {
	var s cofservice.CreateService
	if !bind(c, &s) {
		return
	}
	cof, code := s.Create()
	c.JSON(code, cof)
}

func ConferenceUpdate(c *gin.Context) {
	var s cofservice.UpdateService
	if !bind(c, &s) {
		return
	}
	cof, code := s.Update()
	c.JSON(code, cof)
}

func ConferenceShow(c *gin.Context) {
	cof, code := cofservice.Show(c.Param("id"))
	c.JSON(code, cof)
}

func ConferenceList(c *gin.Context) {
	cofs, code := cofservice.List()
	c.JSON(code, cofs)
}
