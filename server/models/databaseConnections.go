package models

import (
	"database/sql"
	"time"


	_ "github.com/go-sql-driver/mysql"
)

var MySQLClient *sql.DB

func MysqlDBConnect() {
	var err error
	time.Sleep(9 *time.Second)
	MySQLClient, err = sql.Open("mysql", "root:root@tcp(host.docker.internal:20000)/TradeOut")
	if err != nil {
		panic(err)
	}
	err = MySQLClient.Ping()
	if err != nil {
		panic(err)
	}
}
