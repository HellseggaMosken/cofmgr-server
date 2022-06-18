package api

import (
	"cofmgr/service/refereeservice"

	"github.com/gin-gonic/gin"
)

func RefereeAddForConference(c *gin.Context) {
	code := refereeservice.AddRefereeForConference(c.Param("cof_id"), c.Param("user_id"))
	c.Status(code)
}

func RefereeRemoveForConference(c *gin.Context) {
	code := refereeservice.RemoveRefereeForConference(c.Param("cof_id"), c.Param("user_id"))
	c.Status(code)
}

func RefereeAddForContribution(c *gin.Context) {
	code := refereeservice.AddRefereeForContribution(c.Param("ctb_id"), c.Param("user_id"))
	c.Status(code)
}

func RefereeRemoveForContribution(c *gin.Context) {
	code := refereeservice.RemoveRefereeForContribution(c.Param("ctb_id"), c.Param("user_id"))
	c.Status(code)
}
