package connection

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=tiny.db.elephantsql.com user=nawvdgtv password=5tS41gv1lbNCX6Oows-1tjaxhAmKbw1P dbname=nawvdgtv port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate()
	DB = db
}
