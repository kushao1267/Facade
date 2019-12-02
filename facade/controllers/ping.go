package controllers

import "github.com/gin-gonic/gin"

// PingController ...
type PingController struct{}

// Ping test whether if the API server is running
func (ctrl PingController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
