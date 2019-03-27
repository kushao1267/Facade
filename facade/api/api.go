package api

import "github.com/gin-gonic/gin"

func ApiServer(addr ...string) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(addr...) // listen and serve on 0.0.0.0:8080
}