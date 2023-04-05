package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"test-bookingtogo/models"
)

var DB *gorm.DB

func InitDB() {
	var appConfig map[string]string
	appConfig, err := godotenv.Read()

	if err != nil {
		fmt.Println("Error reading .env file")
	}

	postgresCredentials := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		appConfig["DB_HOST"],
		appConfig["DB_USER"],
		appConfig["DB_PASSWORD"],
		appConfig["DB_NAME"],
		appConfig["DB_PORT"],
	)

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: postgresCredentials,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.Nationalitys{})
	DB.AutoMigrate(&models.Customers{})
	DB.AutoMigrate(&models.FamilyLists{})

	DB.Where("1 = 1").Delete(&models.Nationalitys{})
	DB.Where("1 = 1").Delete(&models.Customers{})
	DB.Where("1 = 1").Delete(&models.FamilyLists{})
}
