package api

import (
	"github.com/gin-gonic/gin"
)

// InitRouter: API router
func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", Ping)

	router.POST("/link_preview", LinkPreview)

	return router
}
