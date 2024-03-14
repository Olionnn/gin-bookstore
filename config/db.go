package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func Connection() *sql.DB {
	dataSourceName := viper.GetString("mysql.user") + ":" + viper.GetString("mysql.password") + "@tcp(" + viper.GetString("mysql.host") + ":" + viper.GetString("mysql.port") + ")/" + viper.GetString("mysql.database")
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}


	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db
	// res, _ := db.Query("SELECT * FROM users")
	// for res.Next() {
	// 	var user Users
	// 	err = res.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	fmt.Printf(user.Username)
	// }

}
