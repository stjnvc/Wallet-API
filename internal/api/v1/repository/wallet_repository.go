package repository

import (
	"github.com/stjnvc/wallet-api/internal/api/v1/model"
	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{db}
}

func (wr *WalletRepository) GetWalletByID(id uint) (*model.Wallet, error) {
	var wallet model.Wallet
	if err := wr.db.First(&wallet, id).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (wr *WalletRepository) UpdateWallet(wallet *model.Wallet) error {
	return wr.db.Save(wallet).Error
}
