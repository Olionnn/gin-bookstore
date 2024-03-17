package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var dataSourceName string

var DB *gorm.DB
var err error

func ConnectGorm() error {
	dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.database"))
	DB, err = gorm.Open("mysql", dataSourceName)
	fmt.Println(dataSourceName)
	if err != nil {
		fmt.Println("db err: ", err)
		return err
	}

	fmt.Println("db connected: mysql")
	return nil
}

func Connection() *sql.DB {
	dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.database"))
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db
}
