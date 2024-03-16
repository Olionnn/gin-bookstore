package main

import (
	"fmt"

	"github.com/Olionnn/gin-bookstore/config"
	"github.com/Olionnn/gin-bookstore/http/models"
	"github.com/Olionnn/gin-bookstore/http/routes"
)

func main() {
	config.InitConfig()
	DB, errDB := config.ConnectGorm()
	if errDB != nil {
		panic(errDB)
	}

	errMigrate := DB.AutoMigrate(&models.Users{})
	if errMigrate != nil {
		fmt.Println(errMigrate)
	}

	router := routes.InitRoutes()
	router.Run(":8080")

	// defer db.Close()

}
