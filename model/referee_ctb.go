package model

type RefereeCtb struct {
	UserID         string `gorm:"primaryKey;autoIncrement:false;index"`
	ContributionID string `gorm:"primaryKey;autoIncrement:false;index"`
}
