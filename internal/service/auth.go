package service

import (
	"github.com/tonoy30/clean-arch/internal/domain"
	"github.com/tonoy30/clean-arch/internal/dto"
	"github.com/tonoy30/clean-arch/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignUp(user *dto.User) (*int64, error)
}
type authService struct {
	dao repository.DAO
}

func NewAuthService(dao repository.DAO) AuthService {
	return &authService{dao}
}

func (a authService) SignUp(user *dto.User) (*int64, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(
		user.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	domainUser := domain.User{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Password:    string(hash),
		PhoneNumber: user.PhoneNumber,
		Role:        domain.USER,
	}
	id, err := a.dao.NewUserRepository().CreateUser(domainUser)
	if err != nil {
		return nil, err
	}
	return id, nil
}
