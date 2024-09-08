package services

import (
	"github.com/dealense7/market-price-go/internal/http/requests/auth"
	"github.com/dealense7/market-price-go/internal/interfaces"
	"github.com/dealense7/market-price-go/internal/models"
	"github.com/dealense7/market-price-go/utils"
)

type UserService struct {
	repo interfaces.UserStore
}

func NewUserService(repo interfaces.UserStore) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetByEmail(email string) (*models.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *UserService) GetById(id int) (*models.User, error) {
	return s.repo.GetById(id)
}

func (s *UserService) Create(data *auth.RegisterUserPayload) error {
	hashed, err := utils.HashPassword(data.Password)

	if err != nil {
		return err
	}

	item := models.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  hashed,
	}

	return s.repo.Create(item)
}
