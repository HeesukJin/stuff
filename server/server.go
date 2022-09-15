package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var MySQLClient = mysqlDBConnect()

func main() {
	router := gin.Default()

	router.POST("/register", registerAccount)
	//router.POST("/login", login)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "f",
		})
	})

	router.POST("/login",login)

	fmt.Printf("Starting server at port \n")
	router.Run()
}