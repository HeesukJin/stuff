package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct{
	Username string `json:"firstname" bson:"firstname"`
	Password string `json:"password" bson:"password"`
}

var MongoDBClient = mongoDBConnect()

func main() {
	router := gin.Default()

	//router.POST("/register", registerAccount)
	//router.POST("/login", login)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "f",
		})
	})

	fmt.Printf("Starting server at port \n")
	router.Run()
}