package api

import "github.com/gin-gonic/gin"
import (
	"os"
)

const (
	SuccessCode = 1
	FailCode    = 0

	TestENV    = "test"
	ReleaseENV = "release"
)

func Server(addr ...string) {
	if os.Getenv("APP_ENV") == ReleaseENV {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.TestMode)
	}

	router := InitRouter()
	router.Run(addr...)
}

