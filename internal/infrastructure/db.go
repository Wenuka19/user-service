package infrastructure

import (
	"fmt"
	"github.com/Wenuka19/user-service/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDatabase() *gorm.DB {
	cfg := config.AppConfig
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	var err error
	var DB *gorm.DB

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	return DB
}

func ValidateSchema(db *gorm.DB) {
	var exists bool
	err := db.Raw("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'users')").Scan(&exists).Error
	if err != nil || !exists {
		log.Fatal("User table missing — have you run atlas migrate apply?")
	}
}
