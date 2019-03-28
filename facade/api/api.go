package api

import "github.com/gin-gonic/gin"
import (
	"net/http"
	"strings"
	"github.com/kushao1267/facade/facade/consts"
	"os"
)

func Server(addr ...string) {
	if os.Getenv("APP_ENV") == consts.ReleaseENV {
		gin.SetMode(gin.ReleaseMode)
	}else{
		gin.SetMode(gin.TestMode)
	}

	router := InitRouter()
	router.Run(addr...)
}

// InitRouter: API router
func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", Ping)

	router.POST("/link_preview", LinkPreview)

	return router
}

// Ping: test whether if the API server is running
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// LinkPreview: link preview API
func LinkPreview(c *gin.Context) {
	url := c.Request.FormValue("url")

	if strings.HasPrefix(url, "https") || strings.HasPrefix(url, "http"){
		c.JSON(http.StatusOK, gin.H{
			"code": consts.APISuccessCode,
			"msg": "success",
			"data": url,
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"code": consts.APIFailCode,
		"msg": "fail",
		"data": url,
	})
}
