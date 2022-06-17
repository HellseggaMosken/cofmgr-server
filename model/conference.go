package model

type Conference struct {
	ID          string `gorm:"primaryKey;->;<-:create"` // allow read and create
	Name        string
	StartDate   int64
	EndDate     int64
	URL         string
	Location    string
	Description string
}
