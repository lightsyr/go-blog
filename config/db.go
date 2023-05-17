package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func GetDatabaseConnection() {

	var err error

	DB, err = sql.Open("mysql", "root:admin@/posts?parseTime=true")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database connection done! ✔️")
	}

	err = DB.Ping()

	if err != nil {
		fmt.Println(err.Error())
	}

}
