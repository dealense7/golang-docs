package v1

import (
	"github.com/dealense7/market-price-go/routes/v1/core"
	"github.com/gin-gonic/gin"
)

func RegisterV1Routes(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	core.AuthRoutes(v1)
}
