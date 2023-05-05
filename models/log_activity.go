package models

import "time"

type Activity struct {
	Id   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type LogActivity struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	UserId     int       `gorm:"foreignKey:Id" json:"user_id"`
	ActivityID int       `gorm:"foreignKey:Id" json:"activity_id"`
	CreatedAt  time.Time `gorm:"type:timestamp without time zone;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}
