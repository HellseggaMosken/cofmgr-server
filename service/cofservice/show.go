package cofservice

import (
	"cofmgr/model"
	"cofmgr/service"
	"cofmgr/service/serializer"
)

func List() ([]*serializer.Conference, service.Status) {
	cofs := model.ListConferences()
	return serializer.BuildConferenceres(cofs), service.StatusOK
}

func Show(id string) (*serializer.Conference, service.Status) {
	cof := model.FindConference(id)
	if cof == nil {
		return nil, service.StatusNotFoundItem
	}
	return serializer.BuildConferencere(cof), service.StatusOK
}
