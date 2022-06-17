package refereeservice

import (
	"cofmgr/model"
	"cofmgr/service"
)

func AddRefereeForContribution(contributionID string, userID string) service.Status {
	r := &model.RefereeCtb{
		ContributionID: contributionID,
		UserID:         userID,
	}
	r.Create()
	return service.StatusOK
}

func RemoveRefereeForContribution(contributionID string, userID string) service.Status {
	r := &model.RefereeCtb{
		ContributionID: contributionID,
		UserID:         userID,
	}
	r.Delete()
	return service.StatusOK
}
