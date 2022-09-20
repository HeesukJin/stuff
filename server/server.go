package main

import (
	"fmt"
	"tradeout-server/models"
	"tradeout-server/routes"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
  	"github.com/gin-contrib/sessions/redis"
)

const userkey = "user"

var secret = []byte("secret")

func main() {
	models.MySQLDBConnect()

	router := gin.Default()

	router.POST("/register", routes.RegisterAccount)
	router.POST("/login", routes.Login)
	// Private group, require authentication to access
	private := router.Group("/private")
	private.Use(routes.AuthRequired)
	{
		private.GET("/me", routes.Me)
		private.GET("/status", routes.Status)
	}


	store, _ := redis.NewStore(10, "tcp", "host.docker.internal:9090", "", []byte("secret"))
	store.Options(sessions.Options{MaxAge: 10})

	router.Use(sessions.Sessions("mysession", store))

	router.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}

		if count == 10 {
			session.Clear()
		} else {
			session.Set("count", count)
		}
		session.Save() 

		
		c.JSON(200, gin.H{"count": count})
	})

	fmt.Printf("Starting server at port \n")
	router.Run()
}
