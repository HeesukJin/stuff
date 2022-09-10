package main

import (
	"fmt"
	"context"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
)

func registerAccount(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	user := bson.D{{"username", username}, {"password", hashedPassword}}

	usersCollection := MongoDBClient.Database("TradeOut").Collection("users")
	result, err := usersCollection.InsertOne(context.TODO(), user)

	if err != nil {
        panic(err)
	}
	fmt.Println(result.InsertedID)

	c.JSON(200, gin.H{
		"message": username + " " + password,
	})
}

func login(c *gin.Context){
	var dbUser User

	username := c.PostForm("username")
	password := c.PostForm("password")

	usersCollection := MongoDBClient.Database("TradeOut").Collection("users")
	err:= usersCollection.FindOne(context.TODO(), bson.M{"username":username}).Decode(&dbUser)

	if err != nil {
		c.JSON(401, gin.H{
			"message": "no user",
		})
		return
	}

	fmt.Println(dbUser.Password)

	dbPass:= []byte(dbUser.Password)

	passErr:= bcrypt.CompareHashAndPassword(dbPass, []byte(password))

  	if passErr != nil{
		c.JSON(401, gin.H{
			"message": "bad pass",
		})
		return
  	}

	c.JSON(200, gin.H{
		"message": "f",
	})
}