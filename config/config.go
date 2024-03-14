package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {

	if _, err := os.Stat("config.yaml"); os.IsNotExist(err) {
		fmt.Printf("config.yaml not found\n")
		yml, ymlErr := os.Create("config.yaml")
		if ymlErr != nil {
			log.Fatal(ymlErr)
		}

		defer yml.Close()
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.SetDefault("mysql.port", "")
	viper.SetDefault("mysql.host", "")
	viper.SetDefault("mysql.user", "")
	viper.SetDefault("mysql.password", "")
	viper.SetDefault("mysql.database", "")

	viper.SetDefault("server.port", "")

	if err := viper.WriteConfig(); err != nil {
		fmt.Println("Saving Configuration")
		return
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

}
