package main

import (
	"fmt"
<<<<<<< HEAD
	"os"

	"github.com/kushao1267/Facade/facade/api"
	_ "github.com/kushao1267/Facade/facade/db"
)

var githash = ""
var buildstamp = ""
var goversion = ""

func main() {
	// 二进制文件信息
	args := os.Args
	if len(args) == 2 && (args[1] == "--version" || args[1] == "-v") {
		fmt.Printf("Git Commit Hash: %s\n", githash)
		fmt.Printf("UTC Build Time : %s\n", buildstamp)
		fmt.Printf("Golang Version : %s\n", goversion)
		return
	}
	// server服务
	api.Server("0.0.0.0:8080")
=======
	"github.com/gin-gonic/gin"
	"github.com/kushao1267/Facade/facade/controllers"
	"github.com/kushao1267/Facade/facade/db"
	"os"
)

//CORSMiddleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.New()
	gin.SetMode(os.Getenv("GIN_MODE"))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())

	db.Init()

	v1 := r.Group("/api/v1")
	{
		ping := new(controllers.PingController)
		v1.GET("/ping", ping.Ping)

		link := new(controllers.LinkController)
		v1.POST("/preview", link.Preview)
		v1.POST("/del", link.Del)
	}

	_ = r.Run(":8080")
>>>>>>> 41768819d1a05f2607694f2f4665b9db54b88748
}
