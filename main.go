package main

import (
	"test-bookingtogo/config"
	"test-bookingtogo/routes"
)

func main() {
	config.InitDB()
	routes.NewRouter()
}
