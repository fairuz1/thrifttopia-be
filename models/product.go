package models

import "time"

type Product struct {
	Id             int       `gorm:"primaryKey" json:"id"`
	UserId         int       `gorm:"foreignKey:Id" json:"user_id"`
	CategoryId     int       `gorm:"foreignKey:CategoryIdId;references:Id" json:"category_id"`
	LocationId     int       `gorm:"foreignKey:Id" json:"location_id"`
	PricingId      int       `gorm:"foreignKey:PricingId" json:"pricing_id"`
	Title          string    `gorm:"type:varchar(100)" json:"title"`
	Description    string    `gorm:"type:varchar(255)" json:"description"`
	Images         []Image   `gorm:"-" json:"images"`
	Price          int       `gorm:"type:integer" json:"price"`
	ProofOfPayment string    `gorm:"type:varchar(255)" json:"proof_of_payment"`
	Status         string    `gorm:"type:boolean;not null;default:on_review" json:"status"`
	IsSold         bool      `gorm:"type:boolean;not null;default:false" json:"is_sold"`
	CreatedAt      time.Time `gorm:"type:timestamp without time zone;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"type:timestamp without time zone;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`

	Category Category `gorm:"foreignKey:CategoryId;references:Id" json:"category"`
	Pricing  Pricing  `gorm:"foreignKey:PricingId;references:ID" json:"pricing"`
}

type Category struct {
	Id   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

func (r *Category) TableName() string {
	return "product_categories"
}

type Pricing struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(50)" json:"name"`
	Price       int    `gorm:"type:integer" json:"price"`
	AdsDuration string `gorm:"type:varchar(25)" json:"ads_duration"`
}
