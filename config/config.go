package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"test-bookingtogo/lib/utils"
	"test-bookingtogo/models"
)

var DB *gorm.DB

func InitDB() {

	postgresCredentials := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		utils.GetEnv("DB_HOST"),
		utils.GetEnv("DB_USER"),
		utils.GetEnv("DB_PASSWORD"),
		utils.GetEnv("DB_NAME"),
		utils.GetEnv("DB_PORT"),
	)
	DB, _ = gorm.Open(postgres.New(postgres.Config{
		DSN: postgresCredentials,
	}), &gorm.Config{})
	fmt.Println("Server Success")

	InitMigrate()
	InitSeed()
}

func InitMigrate() {
	DB.AutoMigrate(&models.Nationalities{})
	DB.AutoMigrate(&models.Customers{})
	DB.AutoMigrate(&models.FamilyLists{})

	DB.Where("1 = 1").Delete(&models.Nationalities{})
	DB.Where("1 = 1").Delete(&models.Customers{})
	DB.Where("1 = 1").Delete(&models.FamilyLists{})
}

func InitSeed() error {

	nationalities := []models.Nationalities{
		{NationalityName: "Indonesia", NationalityCode: "62"},
		{NationalityName: "America", NationalityCode: "01"},
		{NationalityName: "Jamaica", NationalityCode: "03"},
	}

	for _, nationality := range nationalities {
		DB.Create(&nationality)
	}

	return nil
}
