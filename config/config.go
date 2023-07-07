package config

import (
	"fmt"
	"go-fiber-app/models"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func getDBConfig() string {

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	DNS := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	return DNS
}

// const DNS = "root:@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

func InitialMigration() {

	DNS := getDBConfig()

	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to Database")
	}
	fmt.Println("Connected to the database")
	DB.AutoMigrate(&models.User{})
}
