package model

import "cofmgr/logger"

type RefereeCtb struct {
	UserID         string `gorm:"primaryKey;autoIncrement:false;index"`
	ContributionID string `gorm:"primaryKey;autoIncrement:false;index"`
}

func (r *RefereeCtb) Create() {
	if err := db.Create(r).Error; err != nil {
		logger.Panic("when create RefereeCtb:", r, err)
	}
}

func (r *RefereeCtb) Delete() {
	if err := db.Delete(r).Error; err != nil {
		logger.Panic("when delete RefereeCtb:", r, err)
	}
}
