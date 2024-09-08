package interfaces

import (
	"github.com/dealense7/market-price-go/internal/http/requests/auth"
	"github.com/dealense7/market-price-go/internal/models"
)

type UserStore interface {
	GetByEmail(email string) (*models.User, error)
	GetById(id int) (*models.User, error)
	Create(user models.User) error
}

type UserService interface {
	GetByEmail(email string) (*models.User, error)
	GetById(id int) (*models.User, error)
	Create(user *auth.RegisterUserPayload) error
}
