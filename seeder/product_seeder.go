package seeder

import (
	"thriftopia/models"

	"gorm.io/gorm"
)

func SeedProduct(db *gorm.DB) {
	products := []models.Product{
		{UserId: 31, CategoryId: 1001, LocationId: 1001, PricingId: 11, Title: "Radio 1990", Description: "Radio jadul tahun 1990", Price: 50000, ProofOfPayment: "https://drive.google.com/file/d/1ymVDbsHP6eWRzSxvKFNykociaujrf076/view?usp=sharing", Status: "rejected", IsSold: false},
		{UserId: 32, CategoryId: 1003, LocationId: 1001, PricingId: 11, Title: "Jaket Kulit", Description: "lorem ipsum lorem ipsum", Price: 180000, ProofOfPayment: "https://drive.google.com/file/d/1ymVDbsHP6eWRzSxvKFNykociaujrf076/view?usp=sharing", Status: "published", IsSold: false},
		{UserId: 31, CategoryId: 1003, LocationId: 1001, PricingId: 10, Title: "Sepatu Olahraga", Description: "lorem ipsum lorem ipsum", Price: 350000, ProofOfPayment: "https://drive.google.com/file/d/1ymVDbsHP6eWRzSxvKFNykociaujrf076/view?usp=sharing", Status: "published", IsSold: false},
		{UserId: 32, CategoryId: 1003, LocationId: 1001, PricingId: 11, Title: "Gitar Elektrik", Description: "lorem ipsum lorem ipsum", Price: 640000, ProofOfPayment: "https://drive.google.com/file/d/1ymVDbsHP6eWRzSxvKFNykociaujrf076/view?usp=sharing", Status: "published", IsSold: false},
		{UserId: 32, CategoryId: 1003, LocationId: 1001, PricingId: 11, Title: "Skintific Moisturizer", Description: "lorem ipsum lorem ipsum", Price: 50000, ProofOfPayment: "https://drive.google.com/file/d/1ymVDbsHP6eWRzSxvKFNykociaujrf076/view?usp=sharing", Status: "published", IsSold: false},
		{UserId: 31, CategoryId: 1001, LocationId: 1001, PricingId: 11, Title: "Radio 1990", Description: "Radio jadul tahun 1990", Price: 50000, ProofOfPayment: "https://drive.google.com/file/d/1ymVDbsHP6eWRzSxvKFNykociaujrf076/view?usp=sharing", Status: "published", IsSold: false},
		{UserId: 32, CategoryId: 1001, LocationId: 1001, PricingId: 11, Title: "Radio 1990", Description: "Radio jadul tahun 1990", Price: 50000, ProofOfPayment: "https://drive.google.com/file/d/1ymVDbsHP6eWRzSxvKFNykociaujrf076/view?usp=sharing", Status: "published", IsSold: false},
		{UserId: 32, CategoryId: 1001, LocationId: 1001, PricingId: 11, Title: "Radio 1990", Description: "Radio jadul tahun 1990", Price: 50000, ProofOfPayment: "https://drive.google.com/file/d/1ymVDbsHP6eWRzSxvKFNykociaujrf076/view?usp=sharing", Status: "published", IsSold: false},
		{UserId: 31, CategoryId: 1001, LocationId: 1001, PricingId: 11, Title: "Radio 1990", Description: "Radio jadul tahun 1990", Price: 50000, ProofOfPayment: "https://drive.google.com/file/d/1ymVDbsHP6eWRzSxvKFNykociaujrf076/view?usp=sharing", Status: "published", IsSold: false},
	}

	for _, product := range products {
		db.Create(&product)
	}

}
