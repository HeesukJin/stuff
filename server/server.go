package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Env struct {
	users interface {
		register() ()
	}
}


func main() {
	router := gin.Default()

	mysql := mysqlDBConnect()

	env := &Env{
		users: BookModel{DB: mysql},
	}

	router.POST("/register", registerAccount)
	//router.POST("/login", login)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "f",
		})
	})

	fmt.Printf("Starting server at port \n")
	router.Run()
}