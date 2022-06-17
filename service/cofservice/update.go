package cofservice

import (
	"cofmgr/model"
	"cofmgr/service"
	"cofmgr/service/serializer"
)

type UpdateService struct {
	ID          string `json:"id"`
	Name        string `json:"name"        binding:"required,min=1,max=100"`
	StartDate   int64  `json:"startDate"   binding:"required,min=0"`
	EndDate     int64  `json:"endDate"     binding:"required,min=0"`
	URL         string `json:"url"         binding:"required,min=1,max=100"`
	Location    string `json:"location"    binding:"max=100"`
	Description string `json:"description" binding:"max=100"`
}

func (s *UpdateService) Update() (*serializer.Conference, service.Status) {
	cof := model.FindConference(s.ID)
	if cof == nil {
		return nil, service.StatusNotFoundItem
	}

	cof.Name = s.Name
	cof.StartDate = s.StartDate
	cof.EndDate = s.EndDate
	cof.URL = s.URL
	cof.Location = s.Location
	cof.Description = s.Description

	cof.Update()

	return serializer.BuildConferencere(cof), service.StatusOK
}
