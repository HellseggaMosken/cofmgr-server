package model

type Contribution struct {
	ID           string `gorm:"primaryKey;->;<-:create"` // allow read and create
	UserID       string `gorm:"index"`
	ConferenceID string `gorm:"index"`
	UpdateDate   int64  `gorm:"autoUpdateTime:milli;<-:update"` // allow read and update
	Title        string
	Abstract     string
	Status       ContributionStatus
	Comment      string
	FileName     string
}

type ContributionStatus int

const (
	ContributionStatusNotAssigned      = -1
	ContributionStatusPending          = 0
	ContributionStatusAccepted         = 1
	ContributionStatusRejected         = 2
	ContributionStatusRevoked          = 3
	ContributionStatusNeedModification = 4
)
