package api

import (
	"cofmgr/service/refereeservice"

	"github.com/gin-gonic/gin"
)

// @Tags Referee
// @Summary Add a Referee for a Conference
// @Description Add a referee(user) for a conference
// @Router /referees/conference/{conference_id}/user/{user_id} [post]
// @Param token header string true "Admin Token"
// @Produce json
// @Success 200 {string} string "Sucess"
// @Failure 403 {string} string "No token or token invalid or expired"
func RefereeAddForConference(c *gin.Context) {
	code := refereeservice.AddRefereeForConference(c.Param("cof_id"), c.Param("user_id"))
	c.Status(code)
}

// @Tags Referee
// @Summary Remove a Referee for a Conference
// @Description Remove a referee(user) for a conference
// @Router /referees/conference/{conference_id}/user/{user_id} [delete]
// @Param token header string true "Admin Token"
// @Produce json
// @Success 200 {string} string "Sucess"
// @Failure 403 {string} string "No token or token invalid or expired"
func RefereeRemoveForConference(c *gin.Context) {
	code := refereeservice.RemoveRefereeForConference(c.Param("cof_id"), c.Param("user_id"))
	c.Status(code)
}

// @Tags Referee
// @Summary Assign a Contribution to a Referee
// @Description Assign a contribution to a referee(user)
// @Router /referees/contribution/{contribution_id}/user/{user_id} [post]
// @Param token header string true "Admin Token"
// @Produce json
// @Success 200 {string} string "Sucess"
// @Failure 403 {string} string "No token or token invalid or expired"
func RefereeAddForContribution(c *gin.Context) {
	code := refereeservice.AddRefereeForContribution(c.Param("ctb_id"), c.Param("user_id"))
	c.Status(code)
}

// @Tags Referee
// @Summary Remove a Contribution for a Referee
// @Description Remove a contribution for a referee(user)
// @Router /referees/contribution/{contribution_id}/user/{user_id} [delete]
// @Param token header string true "Admin Token"
// @Produce json
// @Success 200 {string} string "Sucess"
// @Failure 403 {string} string "No token or token invalid or expired"
func RefereeRemoveForContribution(c *gin.Context) {
	code := refereeservice.RemoveRefereeForContribution(c.Param("ctb_id"), c.Param("user_id"))
	c.Status(code)
}

// @Tags Referee
// @Summary List Conferences for a Referee
// @Description List conferences for a referee
// @Router /referees/conferences/{referee_user_id} [get]
// @Param token header string true "Admin Token"
// @Produce json
// @Success 200 {array}  serializer.Conference "Sucess"
// @Failure 403 {string} string                "No token or token invalid or expired"
func RefereeConferenceList(c *gin.Context) {
	cofs, code := refereeservice.ListConferencesForReferee(c.Param("referee_user_id"))
	c.JSON(code, cofs)
}

// @Tags Referee
// @Summary List Conferences for Current Referee
// @Description List conferences for current referee
// @Router /referees/conferences [get]
// @Param token header string true "user Token"
// @Produce json
// @Success 200 {array}  serializer.Conference "Sucess"
// @Failure 403 {string} string                "No token or token invalid or expired"
func RefereeConferenceListForCurrent(c *gin.Context) {
	cofs, code := refereeservice.ListConferencesForReferee(currentUser(c).ID)
	c.JSON(code, cofs)
}

// @Tags Referee
// @Summary List Contributions for a Referee
// @Description List contributions for a referee
// @Router /referees/contributions/{referee_user_id} [get]
// @Param token header string true "Admin Token"
// @Produce json
// @Success 200 {array}  serializer.Contribution "Sucess"
// @Failure 403 {string} string                  "No token or token invalid or expired"
func RefereeContributionList(c *gin.Context) {
	ctbs, code := refereeservice.ListContributionsForReferee(c.Param("referee_user_id"))
	c.JSON(code, ctbs)
}

// @Tags Referee
// @Summary List Contributions for Current Referee
// @Description List contributions for current referee
// @Router /referees/contributions [get]
// @Param token header string true "User Token"
// @Produce json
// @Success 200 {array}  serializer.Contribution "Sucess"
// @Failure 403 {string} string                  "No token or token invalid or expired"
func RefereeContributionListForCurrent(c *gin.Context) {
	ctbs, code := refereeservice.ListContributionsForReferee(currentUser(c).ID)
	c.JSON(code, ctbs)
}
