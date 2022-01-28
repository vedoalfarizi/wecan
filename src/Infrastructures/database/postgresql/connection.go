package postgresql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/vedoalfarizi/wecan/src/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=wecan password=1sampai8 sslmode=disable TimeZone=Asia/Jakarta")
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.Fundraiser{})

	DB = database
}
