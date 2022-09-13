package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct{
	db *sql.DB
}

func mongoDBConnect() *Database {
	mysqlDB, err := sql.Open("mysql", "root:root@tcp(localhost:20000)/")

    if err != nil {
        panic(err)
    }

	return &Database{
		db: mysqlDB,
	}
}