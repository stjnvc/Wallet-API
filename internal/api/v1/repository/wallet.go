package repository

import (
	"github.com/stjnvc/wallet-api/internal/api/v1/model"
	"gorm.io/gorm"
)

type WalletRepository interface {
	GetWalletByID(walletID int) (*model.Wallet, error)
	UpdateWalletBalance(wallet *model.Wallet) error
}

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &walletRepository{db}
}

func (r *walletRepository) GetWalletByID(walletID int) (*model.Wallet, error) {
	var wallet model.Wallet
	if err := r.db.First(&wallet, walletID).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *walletRepository) UpdateWalletBalance(wallet *model.Wallet) error {
	return r.db.Save(wallet).Error
}
