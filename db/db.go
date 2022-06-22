package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	StorageDB *sql.DB
)

func initDB() {

	dataSource := "root:root123456@tcp(localhost:3306)/storage"

	var err error

	StorageDB, err = sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}
}