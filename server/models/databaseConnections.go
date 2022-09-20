package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-contrib/sessions"

)

var MySQLClient *sql.DB

func MySQLDBConnect() {
	var err error

	MySQLClient, err = sql.Open("mysql", "root:root@tcp(host.docker.internal:20000)/TradeOut")
	if err != nil {
		panic(err)
	}
	
	err = MySQLClient.Ping()
	if err != nil {
		panic(err)
	}
}

func RedisConnect() sessions.Store {
	store, _ := redis.NewStore(10, "tcp", "host.docker.internal:9090", "", []byte("secret"))
	store.Options(sessions.Options{MaxAge: 10})

	return store
}