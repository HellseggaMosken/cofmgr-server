package model

import (
	"errors"
	"strings"

	"cofmgr/logger"

	"gorm.io/gorm"
)

type Admin struct {
	User
}

func FindAdminWithID(id string) *Admin {
	var res Admin
	err := db.First(&res, "id = ?", id).Error
	if err == nil {
		return &res
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Panic("when find admin with id:", id, err)
	}
	return nil
}

func FindAdminWithEmail(email string) *Admin {
	email = strings.ToLower(email)
	var res Admin
	err := db.First(&res, "email = ?", email).Error
	if err == nil {
		return &res
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Panic("when find admin with email:", email, err)
	}
	return nil
}

func ExistsAdminEmail(email string) bool {
	email = strings.ToLower(email)
	var count int64
	err := db.Model(&Admin{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		logger.Panic("when count admins with email:", email, err)
	}
	return count > 0
}
