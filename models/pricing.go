package models

type Pricings struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(50)" json:"name"`
	Price    string `gorm:"type:varchar(50)" json:"price"`
	AdsDuration string `gorm:"type:varchar(25)" json:"ads_duration"`
}