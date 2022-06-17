package userservice

import (
	"strings"

	"cofmgr/model"
	"cofmgr/service"
	"cofmgr/service/serializer"
)

type UpdateService struct {
	Email      string `json:"email"       binding:"required,email,min=1,max=30"`
	FirstName  string `json:"firstName"   binding:"required,min=1,max=20"`
	MidName    string `json:"midName"     binding:"min=1,max=20"`
	LastName   string `json:"lastName"    binding:"required,min=1,max=20"`
	City       string `json:"city"        binding:"min=1,max=20"`
	Country    string `json:"country"     binding:"min=1,max=100"`
	Department string `json:"department"  binding:"min=1,max=100"`
}

func (s *UpdateService) Update(user *model.User) (*serializer.User, service.Status) {
	if strings.ToLower(s.Email) != user.Email {
		if model.ExistsUserEmail(s.Email) {
			return nil, service.StatusEmailExists
		}
	}
	user.Email = s.Email
	user.FirstName = s.FirstName
	user.MidName = s.MidName
	user.LastName = s.LastName
	user.City = s.City
	user.Country = s.Country
	user.Department = s.Department

	user.Update()

	return serializer.BuildUser(user), service.StatusOK
}
