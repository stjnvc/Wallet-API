package service

import (
	"errors"
	"github.com/stjnvc/wallet-api/internal/api/v1/repository"
)

type WalletService struct {
	walletRepo repository.WalletRepository
}

func NewWalletService(walletRepo repository.WalletRepository) *WalletService {
	return &WalletService{walletRepo: walletRepo}
}

func (s *WalletService) GetWalletBalance(walletID int) (float64, error) {
	wallet, err := s.walletRepo.GetWalletByID(walletID)
	if err != nil {
		return 0, err
	}
	return wallet.Balance, nil
}

func (s *WalletService) CreditWallet(walletID int, amount float64) error {
	if amount < 0 {
		return errors.New("amount cannot be negative")
	}

	wallet, err := s.walletRepo.GetWalletByID(walletID)
	if err != nil {
		return errors.New("wallet not found")
	}

	wallet.Balance += amount
	return s.walletRepo.UpdateWalletBalance(wallet)
}

func (s *WalletService) DebitWallet(walletID int, amount float64) error {
	if amount < 0 {
		return errors.New("amount cannot be negative")
	}

	wallet, err := s.walletRepo.GetWalletByID(walletID)
	if err != nil {
		return errors.New("wallet not found")
	}

	if wallet.Balance < amount {
		return errors.New("insufficient balance")
	}

	wallet.Balance -= amount

	return s.walletRepo.UpdateWalletBalance(wallet)
}
