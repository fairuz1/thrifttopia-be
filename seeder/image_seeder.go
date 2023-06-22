package seeder

import (
	"thriftopia/models"

	"gorm.io/gorm"
)

func SeedImages(db *gorm.DB) {
	images := []models.Image{
		{ProductID: 51, URL: "https://drive.google.com/file/d/12xsVz-au38ADtJJkTB1w9bYLJguLlf4H/view?usp=drive_link"},
		{ProductID: 52, URL: "https://drive.google.com/file/d/1UbrSEu4vxm0IaPPXrp--BK4WDT5SYC6c/view?usp=drive_link"},
		{ProductID: 52, URL: "https://drive.google.com/file/d/1V_cGnLkiPsd7g840yhLBslLZPfvuR83Y/view?usp=drive_link"},
		{ProductID: 53, URL: "https://drive.google.com/file/d/1HSm1ESCMCg7wNtHlK-P38KtEN3fIVayQ/view?usp=drive_link"},
		{ProductID: 54, URL: "https://drive.google.com/file/d/1R7zcOYzZQPBCbuqe-2-btSWNbbhhTaAJ/view?usp=drive_link"},
		{ProductID: 55, URL: "https://drive.google.com/file/d/1aA6HbWgZg8aQkV1pkSl22gDvcBQtoqTt/view?usp=drive_link"},
		{ProductID: 56, URL: "https://drive.google.com/file/d/12xsVz-au38ADtJJkTB1w9bYLJguLlf4H/view?usp=drive_link"},
		{ProductID: 57, URL: "https://drive.google.com/file/d/12xsVz-au38ADtJJkTB1w9bYLJguLlf4H/view?usp=drive_link"},
		{ProductID: 58, URL: "https://drive.google.com/file/d/12xsVz-au38ADtJJkTB1w9bYLJguLlf4H/view?usp=drive_link"},
		{ProductID: 59, URL: "https://drive.google.com/file/d/12xsVz-au38ADtJJkTB1w9bYLJguLlf4H/view?usp=drive_link"},
	}

	for _, image := range images {
		db.Create(&image)
	}
}
