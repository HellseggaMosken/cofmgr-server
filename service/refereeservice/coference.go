package refereeservice

import (
	"cofmgr/model"
	"cofmgr/service"
	"cofmgr/service/serializer"
)

func ListRefereesForConference(conferenceID string) ([]*serializer.User, service.Status) {
	referees := model.ListRefereesForConference(conferenceID)
	return serializer.BuildUsers(referees), service.StatusOK
}

func ListConferencesForReferee(refereeUserID string) ([]*serializer.Conference, service.Status) {
	cofs := model.ListConferencesForReferee(refereeUserID)
	return serializer.BuildConferenceres(cofs), service.StatusOK
}

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
