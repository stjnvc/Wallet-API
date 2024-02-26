package migration

import (
	"github.com/stjnvc/wallet-api/internal/api/v1/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&model.Wallet{})
}
