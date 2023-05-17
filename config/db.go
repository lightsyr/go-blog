package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func GetDatabaseConnection() {

	var err error

	dsn := os.Getenv("DATABASE_URL")

	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database connection done! ✅")
	}

	err = DB.Ping()

	if err != nil {
		fmt.Println(err.Error())
	}

}
