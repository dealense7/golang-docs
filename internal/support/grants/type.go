package grants

import "github.com/gin-gonic/gin"

type Grant interface {
	Authenticate(c *gin.Context)
}
