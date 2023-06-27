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
	Pricing  Pricings `gorm:"foreignKey:PricingId;references:ID" json:"pricing"`
	User     Seller   `gorm:"-" json:"seller"`
	Location Location `gorm:"foreignKey:LocationId;references:Id" json:"location"`
}

type Seller struct {
	Name     string `gorm:"-" json:"name"`
	WaNumber string `gorm:"-" json:"wa_number"`
}

type Category struct {
	Id   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

func (r *Category) TableName() string {
	return "product_categories"
}

type Location struct {
	Id         int    `gorm:"primaryKey" json:"id"`
	Province   string `json:"province"`
	City       string `json:"city"`
	Street     string `json:"street"`
	PostalCode int    `json:"postal_code"`
}

func (r *Location) TableName() string {
	return "location"
}
