package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DBHost = "127.0.0.1"
	DBPort = "3306"
	DBUser = "root"
	DBName = "cms"
)

func Connect() *sql.DB {
	dbConn := fmt.Sprintf("%s@tcp(%s:%s)/%s", DBUser, DBHost, DBPort, DBName)
	var db, err = sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	return db
}
