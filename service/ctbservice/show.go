package ctbservice

import (
	"cofmgr/model"
	"cofmgr/service"
	"cofmgr/service/serializer"
)

func Show(id string) (*serializer.Contribution, service.Status) {
	ctb := model.FindContribution(id)
	if ctb == nil {
		return nil, service.StatusNotFoundItem
	}
	return serializer.BuildContribution(ctb), service.StatusOK
}

func ListWithUser(user *model.User) ([]*serializer.Contribution, service.Status) {
	ctbs := model.ListContributionsWithUserID(user.ID)
	return serializer.BuildContributions(ctbs), service.StatusOK
}

func ListWithConference(conferenceID string) ([]*serializer.Contribution, service.Status) {
	ctbs := model.ListContributionsWithConferenceID(conferenceID)
	return serializer.BuildContributions(ctbs), service.StatusOK
}
