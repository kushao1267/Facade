package main

import (
	"fmt"
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
}
