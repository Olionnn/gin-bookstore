package main

import (
	"github.com/Olionnn/gin-bookstore/config"
	"github.com/Olionnn/gin-bookstore/http/routes"
)

func main() {
	config.InitConfig()
	db := config.Connection()

	router := routes.InitRoutes()
	router.Run(":8080")

	defer db.Close()

}
