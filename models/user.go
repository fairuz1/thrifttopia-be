package models

import "time"

type Role struct {
	Id   int  `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type User struct {
	Id        int     `gorm:"primaryKey" json:"id"`
	RoleId    int     `gorm:"foreignKey:Id" json:"role_id"`
	Name      string    `gorm:"type:varchar(300)" json:"name"`
	Email     string    `gorm:"type:varchar(300)" json:"email"`
	Password  string    `gorm:"type:varchar(300)" json:"password"`
	WaNumber  string    `gorm:"type:varchar(300)" json:"wa_number"`
	CreatedAt time.Time `gorm:"type:timestamp without time zone;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp without time zone;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
