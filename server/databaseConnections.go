package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func mysqlDBConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:20000)/")

    if err != nil {
        panic(err)
    }

	return db
}