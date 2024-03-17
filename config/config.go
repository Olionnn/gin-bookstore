package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	if _, err := os.Stat("config.yaml"); os.IsNotExist(err) {
		fmt.Println("config.yaml not found. Creating new config file...")

		yml, ymlErr := os.Create("config.yaml")
		if ymlErr != nil {
			log.Fatal(ymlErr)
		}
		defer yml.Close()

		viper.SetDefault("mysql.port", "3306")
		viper.SetDefault("mysql.host", "localhost")
		viper.SetDefault("mysql.user", "root")
		viper.SetDefault("mysql.password", "121212")
		viper.SetDefault("mysql.database", "bookstore")

		viper.SetDefault("token.lifespan", "1")
		viper.SetDefault("token.secret", "token_paling_aman_di_dunia")

		viper.SetDefault("server.port", "8080")

		if err := viper.WriteConfigAs("config.yaml"); err != nil {
			fmt.Println("Error saving configuration:", err)
			return
		}

		fmt.Println("New config.yaml file created.")
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		fmt.Println("Using existing config.yaml file.")
	}
}
