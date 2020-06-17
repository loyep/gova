package router

import (
	"github.com/gin-gonic/gin"
)

// InitRouter will init router.
func InitRouter() *gin.Engine {

	r := gin.New()

	return r
}
