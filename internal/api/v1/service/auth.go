package service

import (
	"github.com/stjnvc/wallet-api/internal/api/v1/model"
	"github.com/stjnvc/wallet-api/internal/api/v1/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepository *repository.AuthRepository
}

func NewAuthService(authRepository *repository.AuthRepository) *AuthService {
	return &AuthService{authRepository: authRepository}
}

func (as *AuthService) Login(username, password string) (model.User, error) {
	user, err := as.authRepository.FindByUsername(username)
	if err != nil {
		return model.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (as *AuthService) Register(u *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		Username: u.Username,
		Password: string(hashedPassword),
		Email:    u.Email,
	}

	// Attempt to create the user in the repository
	if err := as.authRepository.Create(user); err != nil {
		return err
	}

	return nil
}
