package model

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	Balance float64 `gorm:"default:0"`
}
