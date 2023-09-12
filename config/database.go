package config

import (
	"fmt"
	"os"

//	"github.com/athunlal/models"
    "feyin/bug-tracker/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBconnect() (*gorm.DB, error) {
	uri := os.Getenv("DBURL")
	DB, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Bug{})
	DB.AutoMigrate(&models.Member{})
	DB.AutoMigrate(&models.Project{})
	DB.AutoMigrate(&models.Note{})
	return DB, nil

}