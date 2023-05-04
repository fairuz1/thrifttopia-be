package models

import "time"

type TransactionHistories struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	ProductID int       `gorm:"foreignKey:Id" json:"product_id"`
	BuyerID    int       `gorm:"foreignKey:Id" json:"buyer_id"`
	CreatedAt time.Time `gorm:"type:timestamp without time zone;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}
