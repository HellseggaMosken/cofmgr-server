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

func ListContributionsForReferee(userID string) []Contribution {
	ctbs := make([]Contribution, 0)
	queryCtbs := db.Model(&RefereeCtb{}).Select("contribution_id").Where("user_id = ?", userID)
	query := db.Raw("SELECT * FROM \"contributions\" WHERE id in (?)", queryCtbs)
	err := query.Scan(&ctbs).Error
	if err != nil {
		logger.Panic("when list contributions for referee:", userID, err)
		return nil
	}
	return ctbs
}
