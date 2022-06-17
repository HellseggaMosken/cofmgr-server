package model

import (
	"errors"

	"cofmgr/logger"

	"gorm.io/gorm"
)

type Conference struct {
	ID          string `gorm:"primaryKey;->;<-:create"` // allow read and create
	Name        string
	StartDate   int64
	EndDate     int64
	URL         string
	Location    string
	Description string
}

func (c *Conference) Create() {
	if err := db.Create(c).Error; err != nil {
		logger.Panic("when create conference:", c, err)
	}
}

func (c *Conference) Update() {
	if err := db.Save(c).Error; err != nil {
		logger.Panic("when update conference:", c, err)
	}
}

func ListConferences() []Conference {
	cofs := make([]Conference, 0)
	err := db.Find(&cofs).Error
	if err != nil {
		logger.Panic("when list conferences:", err)
		return nil
	}
	return cofs
}

func FindConference(id string) *Conference {
	var res Conference
	err := db.First(&res, "id = ?", id).Error
	if err == nil {
		return &res
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Panic("when find conference with id:", id, err)
	}
	return nil
}
