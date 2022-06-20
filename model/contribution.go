package model

import (
	"errors"

	"cofmgr/logger"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Contribution struct {
	ID           string `gorm:"primaryKey;->;<-:create"` // allow read and create
	UserID       string `gorm:"index"`
	ConferenceID string `gorm:"index"`
	UpdateDate   int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"` // allow read and update
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

func (c *Contribution) Create() {
	c.ID = uuid.NewV4().String()
	if err := db.Create(c).Error; err != nil {
		logger.Panic("when create contribution:", c, err)
	}
}

func (c *Contribution) Update() {
	if err := db.Save(c).Error; err != nil {
		logger.Panic("when update contribution:", c, err)
	}
}

func ListContributionsWithUserID(userID string) []Contribution {
	ctbs := make([]Contribution, 0)
	err := db.Find(&ctbs, "user_id = ?", userID).Error
	if err != nil {
		logger.Panic("when list contributions with user id:", userID, err)
		return nil
	}
	return ctbs
}

func ListContributionsWithConferenceID(conferenceID string) []Contribution {
	ctbs := make([]Contribution, 0)
	err := db.Find(&ctbs, "conference_id = ?", conferenceID).Error
	if err != nil {
		logger.Panic("when list contributions with conference id:", conferenceID, err)
		return nil
	}
	return ctbs
}

func FindContribution(id string) *Contribution {
	var res Contribution
	err := db.First(&res, "id = ?", id).Error
	if err == nil {
		return &res
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Panic("when find contribution with id:", id, err)
	}
	return nil
}
