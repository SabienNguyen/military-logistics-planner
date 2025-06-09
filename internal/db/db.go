package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/SabienNguyen/military-logistics-planner/internal/models" // import your models here
)

var DB *gorm.DB

func Init() *gorm.DB {
	// Choose environment: dev defaults to SQLite
	env := os.Getenv("APP_ENV") // e.g. "development", "production"
	useSQLite := env == "development" || env == ""

	// Structured logging (shows queries, useful for dev)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n[gorm] ", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info, // change to Warn or Silent in prod
			Colorful:      true,
		},
	)

	var db *gorm.DB
	var err error

	if useSQLite {
		// Local development mode
		db, err = gorm.Open(sqlite.Open("dev.db"), &gorm.Config{
			Logger: newLogger,
		})
	} else {
		// Production/Postgres
		dsn := os.Getenv("DATABASE_URL")
		if dsn == "" {
			dsn = fmt.Sprintf("host=localhost user=postgres password=postgres dbname=logistics port=5432 sslmode=disable")
		}
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
	}

	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	// 🔁 Auto-migrate all models
	err = db.AutoMigrate(
		&models.User{},
		&models.Unit{},
		&models.Depot{},
		&models.Mission{},
		&models.SupplyRequest{},
		&models.Zone{},
		&models.Resource{},
	)
	if err != nil {
		log.Fatalf("auto migration failed: %v", err)
	}

	// 🧠 Set connection pool settings (for production stability)
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get DB object: %v", err)
	}
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db // optional global
	return db
}
