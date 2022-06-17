package cofservice

import (
	"cofmgr/model"
	"cofmgr/service"
	"cofmgr/service/serializer"
)

type CreateService struct {
	Name        string `json:"name"        binding:"required,min=1,max=100"`
	StartDate   int64  `json:"startDate"   binding:"required,min=0"`
	EndDate     int64  `json:"endDate"     binding:"required,min=0"`
	URL         string `json:"url"         binding:"required,min=1,max=100"`
	Location    string `json:"location"    binding:"max=100"`
	Description string `json:"description" binding:"max=100"`
}

func (s *CreateService) Create() (*serializer.Conference, service.Status) {
	cof := &model.Conference{
		Name:        s.Name,
		StartDate:   s.StartDate,
		EndDate:     s.EndDate,
		URL:         s.URL,
		Location:    s.Location,
		Description: s.Description,
	}

	cof.Create()

	return serializer.BuildConferencere(cof), service.StatusOK
}
