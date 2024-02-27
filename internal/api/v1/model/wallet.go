package model

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	ID      uint    `gorm:"primary_key"`
	Balance float64 `gorm:"default:0"`
}
