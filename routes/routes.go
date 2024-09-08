package routes

import (
	v1 "github.com/dealense7/market-price-go/routes/v1"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() {
	r := gin.Default()
	api := r.Group("/api")
	v1.RegisterV1Routes(api)

	r.Run(":8080")
}
