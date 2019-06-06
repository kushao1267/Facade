package main

import (
	"github.com/kushao1267/Facade/facade/api"
	_ "github.com/kushao1267/Facade/facade/db"
)

func main() {
	api.Server("0.0.0.0:8080")
}
