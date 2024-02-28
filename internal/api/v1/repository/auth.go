package repository

import (
	"github.com/stjnvc/wallet-api/internal/api/v1/model"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (ar *AuthRepository) FindByUsername(username string) (model.User, error) {
	var user model.User
	err := ar.db.Where("username = ?", username).First(&user).Error
	return user, err
}

func (ar *AuthRepository) Create(user *model.User) error {
	err := ar.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
