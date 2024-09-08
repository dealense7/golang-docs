package grants

import (
	"errors"
	"net/http"

	"github.com/dealense7/market-price-go/config"
	"github.com/dealense7/market-price-go/internal/http/requests/auth"
	"github.com/dealense7/market-price-go/internal/interfaces"
	"github.com/dealense7/market-price-go/pkg"
	"github.com/dealense7/market-price-go/utils"
	"github.com/gin-gonic/gin"
)

var ErrWrongCredentials = errors.New("wrong credentials")

type PasswordGrant struct {
	service interfaces.UserService
}

func NewPasswordGrant(service interfaces.UserService) *PasswordGrant {
	return &PasswordGrant{service: service}
}

func (g *PasswordGrant) Authenticate(c *gin.Context) {
	var payload auth.LoginUserPayload

	// Parse JSON
	if err := utils.ParseJson(c, &payload); err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	// Validate
	if errs := utils.Validate(payload); errs != nil {
		utils.WriteJson(c, http.StatusUnprocessableEntity, errs)
		return
	}

	// Check if user exists
	item, err := g.service.GetByEmail(payload.Email)

	if err != nil {
		if errors.Is(err, utils.ErrResourceNotFound) {
			utils.WriteError(c, http.StatusNotFound, err)
			return
		}

		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	// Compare Passwords
	if !utils.ComparePassword(item.Password, payload.Password) {
		utils.WriteError(c, http.StatusBadRequest, ErrWrongCredentials)
		return
	}

	// Generate Token
	token, err := pkg.CreateJwtToken(config.Envs.JWTSecret, item.ID)

	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(c, http.StatusOK, map[string]string{"token": token})
}
