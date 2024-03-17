package main

import (
	"fmt"

	"github.com/Olionnn/gin-bookstore/config"
	"github.com/Olionnn/gin-bookstore/http/models"
	"github.com/Olionnn/gin-bookstore/http/routes"
	"github.com/spf13/viper"
)

func main() {
	config.InitConfig()
	errDB := config.ConnectGorm()
	if errDB != nil {
		panic(errDB)
	}

	errMigrate := config.DB.AutoMigrate(&models.Users{})
	if errMigrate != nil {
		fmt.Println(errMigrate)
	}

	router := routes.InitRoutes()
	router.Run(":" + viper.GetString("server.port"))

	// defer db.Close()

}
