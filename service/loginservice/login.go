package loginservice

import (
	"cofmgr/model"
	"cofmgr/service"
	"cofmgr/service/authservice"
	"cofmgr/service/serializer"
)

type LoginService struct {
	Email    string `json:"email"    binding:"required,email,min=1,max=30"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

func (s *LoginService) LoginUser() (u *serializer.User, token string, st service.Status) {
	user := model.FindUserWithEmail(s.Email)
	if user == nil || !model.PasswordCompare(s.Password, user.PasswordDigest) {
		return nil, "", service.StatusWrongAccount
	}
	token = authservice.CreateToken(user.ID, false)
	u = serializer.BuildUser(user)
	return u, token, service.StatusOK
}

func (s *LoginService) LoginAdmin() (u *serializer.User, token string, st service.Status) {
	admin := model.FindAdminWithEmail(s.Email)
	if admin == nil || !model.PasswordCompare(s.Password, admin.PasswordDigest) {
		return nil, "", service.StatusWrongAccount
	}
	token = authservice.CreateToken(admin.ID, true)
	u = serializer.BuildUser(&admin.User)
	return u, token, service.StatusOK
}
