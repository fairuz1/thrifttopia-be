package seeder

import (
	"thriftopia/models"

	"gorm.io/gorm"
)

func SeedLogAcitvities(db *gorm.DB) {
	log_activities := []models.LogActivity{
		{UserId: 31, ActivityID: 1001},
		{UserId: 32, ActivityID: 1001},
	}

	for _, pricing := range log_activities {
		db.Create(&pricing)
	}
}
