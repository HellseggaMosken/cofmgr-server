package userservice

import (
	"cofmgr/model"
	"cofmgr/service"
	"cofmgr/service/serializer"
)

type RegisterService struct {
	Email      string `json:"email"       binding:"required,email,min=1,max=30"`
	FirstName  string `json:"firstName"   binding:"required,min=1,max=20"`
	MidName    string `json:"midName"     binding:"min=1,max=20"`
	LastName   string `json:"lastName"    binding:"required,min=1,max=20"`
	City       string `json:"city"        binding:"min=1,max=20"`
	Country    string `json:"country"     binding:"min=1,max=100"`
	Department string `json:"department"  binding:"min=1,max=100"`
	Password   string `json:"password"    binding:"required,min=6,max=20"`
}

func (s *RegisterService) Register() (*serializer.User, service.Status) {
	user := &model.User{
		Email:          s.Email,
		FirstName:      s.FirstName,
		MidName:        s.MidName,
		LastName:       s.LastName,
		City:           s.City,
		Country:        s.Country,
		Department:     s.Department,
		PasswordDigest: model.PasswordDigest(s.Password),
	}

	if model.ExistsUserEmail(s.Email) {
		return nil, service.StatusEmailExists
	}

	user.Create()

	return serializer.BuildUser(user), service.StatusOK
}
