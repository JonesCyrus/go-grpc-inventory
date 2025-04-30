package config

import (
	"fmt"

	"go-grpc-inventory/pkg/models"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(vi *viper.Viper) (*gorm.DB, error) {
	USER := vi.GetString("DB_USER")
	HOST := vi.GetString("DB_HOST")
	PORT := vi.GetString("DB_PORT")
	NAME := vi.GetString("DB_NAME")
	PASS := vi.GetString("DB_PASSWORD")

	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", HOST, USER, PASS, NAME, PORT)
	fmt.Println("Connecting to database")
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database: ", err)
		return db, err
	}

	fmt.Println("Connected to database")

	return db, nil
}

func MigrateDatabase(tx *gorm.DB) {
	var err error
	err = tx.AutoMigrate(&models.Inventory{})
	if err != nil {
		panic("Failed to migrate user table to database")
	}
}
