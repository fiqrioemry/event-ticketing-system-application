package config

import (
	"fmt"
	"time"

	"github.com/fiqrioemry/event_ticketing_system_app/server/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {

	// Connect to MySQL server
	dbRoot, err := gorm.Open(mysql.Open(AppConfig.DatabaseRootURL), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to MySQL server: " + err.Error())
	}

	// Create DB if not exists
	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", AppConfig.DatabaseName)
	if err := dbRoot.Exec(sql).Error; err != nil {
		panic("Failed to create database: " + err.Error())
	}

	// Connect to main DB
	for range 10 {
		DB, err = gorm.Open(mysql.Open(AppConfig.DatabaseURL), &gorm.Config{})
		if err == nil {
			break
		}
		fmt.Println("Waiting for database to be ready...")
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// Migration database schema
	if err := DB.AutoMigrate(
		&models.User{},
		&models.Event{},
		&models.Order{},
		&models.Ticket{},
		&models.OrderDetail{},
		&models.Payment{},
		&models.UserTicket{},
		&models.WithdrawalRequest{},
	); err != nil {
		panic("Migration failed: " + err.Error())
	}

	sqlDB, err := DB.DB()
	if err != nil {
		panic("Failed to get database connection: " + err.Error())
	}
	// Set connection pool settings
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("✅ Database configured")
}
