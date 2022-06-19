package model

import "cofmgr/logger"

type RefereeCof struct {
	UserID       string `gorm:"primaryKey;autoIncrement:false;index"`
	ConferenceID string `gorm:"primaryKey;autoIncrement:false;index"`
}

func (r *RefereeCof) Create() {
	if err := db.Create(r).Error; err != nil {
		logger.Panic("when create RefereeCof:", r, err)
	}
}

func (r *RefereeCof) Delete() {
	if err := db.Delete(r).Error; err != nil {
		logger.Panic("when delete RefereeCof:", r, err)
	}
}

func ListRefereesForConference(conferenceID string) []User {
	users := make([]User, 0)
	queryReferees := db.Model(&RefereeCof{}).Select("user_id").Where("conference_id = ?", conferenceID)
	query := db.Raw("SELECT * FROM \"users\" WHERE id in ?", queryReferees)
	err := query.Scan(&users).Error
	if err != nil {
		logger.Panic("when list referees for conference:", conferenceID, err)
		return nil
	}
	return users
}

func ListConferencesForReferee(userID string) []Conference {
	cofs := make([]Conference, 0)
	queryCofs := db.Model(&RefereeCof{}).Select("conference_id").Where("user_id = ?", userID)
	query := db.Raw("SELECT * FROM \"conferences\" WHERE id in ?", queryCofs)
	err := query.Scan(&cofs).Error
	if err != nil {
		logger.Panic("when list conferences for referee:", userID, err)
		return nil
	}
	return cofs
}
