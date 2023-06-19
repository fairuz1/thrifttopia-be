package models

type Image struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	ProductID int    `gorm:"foreignKey:ID" json:"product_id"`
	URL       string `gorm:"type:varchar(255)" json:"url"`
}
