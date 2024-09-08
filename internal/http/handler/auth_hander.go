package handler

import (
	"errors"
	"net/http"

	"github.com/dealense7/market-price-go/internal/http/requests/auth"
	"github.com/dealense7/market-price-go/internal/interfaces"
	"github.com/dealense7/market-price-go/internal/support/grants"
	"github.com/dealense7/market-price-go/utils"
	"github.com/gin-gonic/gin"
)

var (
	ErrUserWithEmailAlreadyExists = errors.New("user with same email already exists")
	ErrWrongCredentials           = errors.New("wrong credentials")
	ErrGrantNotSupported          = errors.New("grant type is not supported")
)

type AuthHandler struct {
	service interfaces.UserService
	grants  map[string]grants.Grant
}

func NewAuthHandler(service interfaces.UserService) *AuthHandler {
	return &AuthHandler{
		service: service,
		grants:  make(map[string]grants.Grant),
	}
}

func (h *AuthHandler) RegisterGrant(name string, grant grants.Grant) {
	h.grants[name] = grant
}

func (h *AuthHandler) Login(c *gin.Context) {
	grantType := c.GetHeader("Grant-Type")

	grant, exists := h.grants[grantType]
	if !exists {
		utils.WriteError(c, http.StatusBadRequest, ErrGrantNotSupported)
		return
	}

	grant.Authenticate(c)
}

func (h *AuthHandler) GetMe(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.WriteServerError(c)
		return
	}

	utils.WriteJson(c, http.StatusOK, user)
}

func (h *AuthHandler) Register(c *gin.Context) {
	var payload auth.RegisterUserPayload

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
	item, err := h.service.GetByEmail(payload.Email)

	if err != nil && !errors.Is(err, utils.ErrResourceNotFound) {
		utils.WriteServerError(c)
		return
	}

	if item != nil {
		utils.WriteError(c, http.StatusBadRequest, ErrUserWithEmailAlreadyExists)
		return
	}

	// Create user
	if err := h.service.Create(&payload); err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	utils.WriteMessage(c, http.StatusCreated, "user successfuly created!")
}
