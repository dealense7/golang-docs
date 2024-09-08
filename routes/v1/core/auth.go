package core

import (
	"github.com/dealense7/market-price-go/internal/http/handler"
	"github.com/dealense7/market-price-go/internal/http/repository/user"
	"github.com/dealense7/market-price-go/internal/http/services"
	"github.com/dealense7/market-price-go/internal/interfaces"
	"github.com/dealense7/market-price-go/internal/middleware"
	"github.com/dealense7/market-price-go/internal/support/grants"
	"github.com/dealense7/market-price-go/utils"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {

	store := user.NewStore(utils.DB)
	service := services.NewUserService(store)
	handler := handler.NewAuthHandler(service)

	registerGrants(handler, service)

	auth := router.Group("auth")
	{
		auth.POST("/login", handler.Login)
		auth.POST("/register", handler.Register)

		auth.Use(middleware.JWTAuthMiddleware(service))
		{
			auth.GET("/me", handler.GetMe)
		}
	}
}

func registerGrants(h *handler.AuthHandler, s interfaces.UserService) {
	passwordGrant := grants.NewPasswordGrant(s)
	h.RegisterGrant("password", passwordGrant)

}
