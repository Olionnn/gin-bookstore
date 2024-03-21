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

	if errMigrate := config.DB.AutoMigrate(&models.Users{}); errMigrate.Error != nil {
		fmt.Println(errMigrate.Error)
	}

	if errMigrate := config.DB.AutoMigrate(&models.Books{}); errMigrate.Error != nil {
		fmt.Println(errMigrate.Error)
	}

	router := routes.InitRoutes()
	router.Run(":" + viper.GetString("server.port"))

	// defer db.Close()

}
