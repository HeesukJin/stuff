package main

import (
	"fmt"
	"tradeout-server/models"
	"tradeout-server/routes"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"github.com/gin-contrib/sessions"
)

const userkey = "user"

var secret = []byte("secret")

func main() {
	router := gin.Default()

	models.MySQLDBConnect()
	router.Use(sessions.Sessions("mysession", models.RedisConnect()))
	//only temporary
	router.Use(cors.Default())


	router.POST("/register", routes.RegisterAccount)
	router.POST("/login", routes.Login)
	// Private group, require authentication to access
	private := router.Group("/private")
	private.Use(routes.AuthRequired)
	{
		private.GET("/me", routes.Me)
		private.GET("/status", routes.Status)
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"count": "count"})
	})

	fmt.Printf("Starting server at port \n")
	router.Run()
}
