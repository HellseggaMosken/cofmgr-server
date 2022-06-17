package model

type RefereeCof struct {
	UserID       string `gorm:"primaryKey;autoIncrement:false;index"`
	ConferenceID string `gorm:"primaryKey;autoIncrement:false;index"`
}
