package passwordservice

import (
	"cofmgr/model"
	"cofmgr/service"
)

type PasswordService struct {
	OldPassword string `json:"oldPassword" binding:"required,min=6,max=20"`
	NewPassword string `json:"newPassword" binding:"required,min=6,max=20"`
}

func (s *PasswordService) ChangeUserPassword(user *model.User) service.Status {
	if model.PasswordCompare(s.OldPassword, user.PasswordDigest) {
		return service.StatusWrongAccount
	}

	user.PasswordDigest = model.PasswordDigest(s.NewPassword)
	user.Update()
	return service.StatusOK
}

func (s *PasswordService) ChangeAdminPassword(admin *model.Admin) service.Status {
	if model.PasswordCompare(s.OldPassword, admin.PasswordDigest) {
		return service.StatusWrongAccount
	}

	admin.PasswordDigest = model.PasswordDigest(s.NewPassword)
	admin.Update()
	return service.StatusOK
}
