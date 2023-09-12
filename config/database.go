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

	/*
		uri := os.Getenv("DBURL")
		DB, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
		if err != nil {
			fmt.Println(err)
		}
	*/

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	// uri := "postgres://bug_tracker_api_user:MFxpttTB0B2hiVFyB2BDKOFYee8soeXb@dpg-ck09mah5mpss73ddmvd0-a/bug_tracker_api"

	uri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
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
