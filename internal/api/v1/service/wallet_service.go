package service

import (
	"errors"
	"github.com/stjnvc/wallet-api/internal/api/v1/repository"
)

var ErrNegativeAmount = errors.New("negative amount is not allowed")
var ErrInsufficientFunds = errors.New("Insufficient funds.")

type WalletService struct {
	walletRepo *repository.WalletRepository
}

func NewWalletService(walletRepo *repository.WalletRepository) *WalletService {
	return &WalletService{walletRepo}
}

func (ws *WalletService) GetBalance(walletID uint) (float64, error) {
	wallet, err := ws.walletRepo.GetWalletByID(walletID)
	if err != nil {
		return 0, err
	}
	return wallet.Balance, nil
}

func (ws *WalletService) Credit(walletID uint, amount float64) error {
	if amount < 0 {
		return ErrNegativeAmount
	}
	wallet, err := ws.walletRepo.GetWalletByID(walletID)
	if err != nil {
		return err
	}
	wallet.Balance += amount
	return ws.walletRepo.UpdateWallet(wallet)
}

func (ws *WalletService) Debit(walletID uint, amount float64) error {
	if amount < 0 {
		return ErrNegativeAmount
	}
	wallet, err := ws.walletRepo.GetWalletByID(walletID)
	if err != nil {
		return err
	}
	if wallet.Balance < amount {
		return ErrInsufficientFunds
	}
	wallet.Balance -= amount
	return ws.walletRepo.UpdateWallet(wallet)
}
