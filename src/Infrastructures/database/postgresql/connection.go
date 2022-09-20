package postgresql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/vedoalfarizi/wecan/src/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("postgres", "host=localhost port=5492 user=root dbname=wecan password=pass sslmode=disable TimeZone=Asia/Jakarta")
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.Fundraiser{})

	DB = database
}
