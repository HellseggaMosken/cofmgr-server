package ctbservice

import (
	"cofmgr/model"
	"cofmgr/service"
	"cofmgr/service/serializer"
)

type CreateService struct {
	UserID       string `json:"userId"       binding:"required"`
	ConferenceID string `json:"conferenceId" binding:"required"`
	Title        string `json:"title"        binding:"required,min=10,max=100"`
	Abstract     string `json:"abstract"     binding:"required,min=10,max=2000"`
	FileName     string `json:"fileName"     binding:"required"`
}

func (s *CreateService) Create() (*serializer.Contribution, service.Status) {
	ctb := &model.Contribution{
		UserID:       s.UserID,
		ConferenceID: s.ConferenceID,
		Title:        s.Title,
		Abstract:     s.Abstract,
		FileName:     s.FileName,
		Status:       model.ContributionStatusNotAssigned,
	}

	ctb.Create()

	return serializer.BuildContribution(ctb), service.StatusOK
}