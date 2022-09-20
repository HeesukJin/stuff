package main

import (
	"fmt"

	"tradeout-server/models"
	"tradeout-server/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const userkey = "user"

var secret = []byte("secret")

func main() {
	models.MysqlDBConnect()

	router := gin.Default()
	store := cookie.NewStore(secret)
	router.Use(sessions.Sessions("mysession", store))

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
		session := sessions.Default(c)
		message := "Bad"
		if session.Get("ID") != nil {
			message = session.Get("name").(string)
		}
		c.JSON(200, gin.H{
			"message": message,
		})
	})

	fmt.Printf("Starting server at port \n")
	router.Run()
}
