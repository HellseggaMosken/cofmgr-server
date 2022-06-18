package api

import (
	"cofmgr/service/refereeservice"

	"github.com/gin-gonic/gin"
)

// @Tags Referee
// @Summary Add a Referee for a Conference
// @Description Add a referee(user) for a conference
// @Router /referees/conference/{conference_id}/user/{user_id} [post]
// @Param token body string true "Admin Token"
// @Produce json
// @Success 200 {string} string "Sucess"
func RefereeAddForConference(c *gin.Context) {
	code := refereeservice.AddRefereeForConference(c.Param("cof_id"), c.Param("user_id"))
	c.Status(code)
}

// @Tags Referee
// @Summary Remove a Referee for a Conference
// @Description Remove a referee(user) for a conference
// @Router /referees/conference/{conference_id}/user/{user_id} [delete]
// @Param token body string true "Admin Token"
// @Produce json
// @Success 200 {string} string "Sucess"
func RefereeRemoveForConference(c *gin.Context) {
	code := refereeservice.RemoveRefereeForConference(c.Param("cof_id"), c.Param("user_id"))
	c.Status(code)
}

// @Tags Referee
// @Summary Assign a Contribution to a Referee
// @Description Assign a contribution to a referee(user)
// @Router /referees/contribution/{contribution_id}/user/{user_id} [post]
// @Param token body string true "Admin Token"
// @Produce json
// @Success 200 {string} string "Sucess"
func RefereeAddForContribution(c *gin.Context) {
	code := refereeservice.AddRefereeForContribution(c.Param("ctb_id"), c.Param("user_id"))
	c.Status(code)
}

// @Tags Referee
// @Summary Remove a Contribution for a Referee
// @Description Remove a contribution for a referee(user)
// @Router /referees/contribution/{contribution_id}/user/{user_id} [delete]
// @Param token body string true "Admin Token"
// @Produce json
// @Success 200 {string} string "Sucess"
func RefereeRemoveForContribution(c *gin.Context) {
	code := refereeservice.RemoveRefereeForContribution(c.Param("ctb_id"), c.Param("user_id"))
	c.Status(code)
}
