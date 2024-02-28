package db

import (
	"github.com/stjnvc/wallet-api/internal/api/v1/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dsn := "root:pass1234@tcp(localhost:3306)/wallet-api-db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db

	return nil
}

func Close() error {
	if DB == nil {
		return nil // Database connection is not initialized
	}
	dbConn, err := DB.DB()
	if err != nil {
		return err
	}
	return dbConn.Close()
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Wallet{},
		&model.User{},
	)
}

func Seed() error {
	// Drop existing table if exists
	if err := DB.Migrator().DropTable(&model.Wallet{}); err != nil {
		return err
	}

	// Auto migrate to create table
	if err := DB.AutoMigrate(&model.Wallet{}); err != nil {
		return err
	}

	// Create sample wallets
	wallets := []model.Wallet{
		{Balance: 100.0},
		{Balance: 200.0},
		{Balance: 300.0},
		{Balance: 400.0},
		{Balance: 500.0},
		{Balance: 600.0},
		{Balance: 700.0},
		{Balance: 800.0},
		{Balance: 900.0},
		{Balance: 1000.0},
	}

	// Insert wallets into the database
	if err := DB.Create(&wallets).Error; err != nil {
		return err
	}

	return nil
}
