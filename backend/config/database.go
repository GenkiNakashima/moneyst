package config

import (
	"fmt"
	"log"
	"os"

	"github.com/GenkiNakashima/moneyst/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		// Fallback to individual environment variables
		host := getEnv("DB_HOST", "localhost")
		port := getEnv("DB_PORT", "5432")
		user := getEnv("DB_USER", "postgres")
		password := getEnv("DB_PASSWORD", "postgres")
		dbname := getEnv("DB_NAME", "moneyst")

		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")

	// Auto Migrate models
	err = DB.AutoMigrate(
		&models.User{},
		&models.Lesson{},
		&models.Curriculum{},
		&models.DailyLesson{},
		&models.SimulationTrade{},
		&models.TradeDiary{},
		&models.SummarizedNews{},
		&models.Post{},
		&models.Reply{},
		&models.PostLike{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
