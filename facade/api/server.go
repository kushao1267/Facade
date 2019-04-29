package api

import (
	"github.com/gin-gonic/gin"
	"os"
)

const (
	SuccessCode = 1
	FailCode    = 0

	TestENV    = "test"
	ReleaseENV = "release"
)

func Server(addr ...string) {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if os.Getenv("GIN_MODE") == ReleaseENV {
		gin.SetMode(gin.ReleaseMode)
	} else if os.Getenv("GIN_MODE") == TestENV {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r.GET("/ping", Ping)
	r.POST("/link_preview", LinkPreview)
	r.POST("/del_cache", DelCache)

	r.Run(addr...)
}
