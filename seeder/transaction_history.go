package seeder

import (
	"thriftopia/models"

	"gorm.io/gorm"
)

func SeedTransactionHistories(db *gorm.DB) {
	transactions := []models.TransactionHistories{
		{ProductID: 50, BuyerID: 31},
	}

	for _, transaction := range transactions {
		db.Create(&transaction)
	}
}
