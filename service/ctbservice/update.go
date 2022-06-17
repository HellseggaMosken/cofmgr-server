package ctbservice

import (
	"cofmgr/model"
	"cofmgr/service"
	"cofmgr/service/serializer"
)

type UpdateService struct {
	ID       string `json:"id"           binding:"required"`
	Title    string `json:"title"        binding:"required,min=10,max=100"`
	Abstract string `json:"abstract"     binding:"required,min=10,max=2000"`
	Status   int    `json:"status"       binding:"required,min=-1,max=4"`
	Comment  string `json:"comment"      binding:"max=2000"`
	FileName string `json:"fileName"     binding:"required"`
}

func (s *UpdateService) Update() (*serializer.Contribution, service.Status) {
	ctb := model.FindContribution(s.ID)
	if ctb == nil {
		return nil, service.StatusNotFoundItem
	}

	ctb.Title = s.Title
	ctb.Abstract = s.Abstract
	ctb.Status = model.ContributionStatus(s.Status)
	ctb.Comment = s.Comment
	ctb.FileName = s.FileName

	ctb.Create()

	return serializer.BuildContribution(ctb), service.StatusOK
}
