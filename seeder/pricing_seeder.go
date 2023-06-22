package seeder

import (
	"thriftopia/models"

	"gorm.io/gorm"
)

func SeedPricing(db *gorm.DB) {
	pricings := []models.Pricings{
		{Name: "Dengan iklan di Dashboard ", Price: 8000, AdsDuration: "7 Days"},
		{Name: "Tanpa iklan ", Price: 4000, AdsDuration: "0 Days"},
	}

	for _, pricing := range pricings {
		db.Create(&pricing)
	}
}
