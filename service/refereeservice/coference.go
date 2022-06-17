package refereeservice

import (
	"cofmgr/model"
	"cofmgr/service"
)

func AddRefereeForConference(conferenceID string, userID string) service.Status {
	r := &model.RefereeCof{
		ConferenceID: conferenceID,
		UserID:       userID,
	}
	r.Create()
	return service.StatusOK
}

func RemoveRefereeForConference(conferenceID string, userID string) service.Status {
	r := &model.RefereeCof{
		ConferenceID: conferenceID,
		UserID:       userID,
	}
	r.Delete()
	return service.StatusOK
}
