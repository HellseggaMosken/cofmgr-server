package refereeservice

import (
	"cofmgr/model"
	"cofmgr/service"
	"cofmgr/service/serializer"
)

func ListContributionsForReferee(refereeUserID string) ([]*serializer.Contribution, service.Status) {
	ctbs := model.ListContributionsForReferee(refereeUserID)
	return serializer.BuildContributions(ctbs), service.StatusOK
}

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

func ListContributionsNotAssignedForConference(conferenceID string) ([]*serializer.Contribution, service.Status) {
	ctbs := model.ListContributionsNotAssignedForConference(conferenceID)
	return serializer.BuildContributions(ctbs), service.StatusOK
}
