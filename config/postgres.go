package config

import (
	"fmt"
	"log"
	"time"

	"github.com/chef-01/live-tracking-server/commons/constants"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitPostgres() {
	constants.LoadEnv()

	host := constants.GetEnv("DB_HOST", "localhost")
	port := constants.GetEnv("DB_PORT", "5432")
	user := constants.GetEnv("DB_USER", "postgres")
	password := constants.GetEnv("DB_PASSWORD", "manab123")
	dbName := constants.GetEnv("DB_NAME", "postgres")
	sslMode := constants.GetEnv("DB_SSLMODE", "disable")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		host, user, password, dbName, port, sslMode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get DB instance: %v", err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("✅ Connected to Postgres")
}
