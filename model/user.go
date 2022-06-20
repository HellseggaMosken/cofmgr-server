package model

import (
	"errors"
	"strings"

	"cofmgr/logger"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID             string `gorm:"primaryKey;->;<-:create"` // allow read and create
	Email          string `gorm:"index;not null;unique"`
	FirstName      string
	MidName        string
	LastName       string
	City           string
	Country        string
	Department     string
	PasswordDigest string `gorm:"not null"`
}

func (u *User) Update() {
	u.Email = strings.ToLower(u.Email)
	if err := db.Save(u).Error; err != nil {
		logger.Panic("when update user:", u, err)
	}
}

func (u *User) Create() {
	u.ID = uuid.NewV4().String()
	u.Email = strings.ToLower(u.Email)
	if err := db.Create(u).Error; err != nil {
		logger.Panic("when create user:", u, err)
	}
}

func ListUsers() []User {
	users := make([]User, 0)
	err := db.Find(&users).Error
	if err != nil {
		logger.Panic("when list users:", err)
		return nil
	}
	return users
}

func FindUserWithID(id string) *User {
	var res User
	err := db.First(&res, "id = ?", id).Error
	if err == nil {
		return &res
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Panic("when find user with id:", id, err)
	}
	return nil
}

func FindUserWithEmail(email string) *User {
	email = strings.ToLower(email)
	var res User
	err := db.First(&res, "email = ?", email).Error
	if err == nil {
		return &res
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Panic("when find user with email:", email, err)
	}
	return nil
}

func ExistsUserEmail(email string) bool {
	email = strings.ToLower(email)
	var count int64
	err := db.Model(&User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		logger.Panic("when count users with email:", email, err)
	}
	return count > 0
}
