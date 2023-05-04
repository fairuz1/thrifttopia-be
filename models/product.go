package models

import "time"

type Product struct {
	Id          int       `gorm:"primaryKey" json:"id"`
	UserId      int       `gorm:"foreignKey:Id" json:"user_id"`
	CategoryId  int       `gorm:"foreignKey:Id" json:"category_id"`
	LocationId  int       `gorm:"foreignKey:Id" json:"location_id"`
	PricingId   int       `gorm:"foreignKey:Id" json:"pricing_id"`
	Title       string    `gorm:"type:varchar(100)" json:"title"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	Images      string    `gorm:"type:varchar(20)" json:"images"`
	IsReviewed  bool      `gorm:"type:boolean;not null;default:false" json:"is_reviewed"`
	IsSold      bool      `gorm:"type:boolean;not null;default:false" json:"is_sold"`
	CreatedAt   time.Time `gorm:"type:timestamp without time zone;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp without time zone;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
