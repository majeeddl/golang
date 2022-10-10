package databases

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func DB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgis@2022 dbname=webapp port=5432 sslmode=disable TimeZone=Asia/Tehran"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting postgres database", err)
	}

	return db
}
