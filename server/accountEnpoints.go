package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

func registerAccount(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	usernameErr := validateUsername(username)
	if usernameErr != nil{
		c.String(401, usernameErr.Error())

		return
	}

	passwordErr := valididatePassword(password)
	if passwordErr != nil{
		c.String(401, passwordErr.Error())

		return
	}

	if usernameExists(username) == true {
		c.String(401, "User already exists.")

		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		c.String(500, "Something went wrong on our end, please try again.")

		return
	}

	user := User{Username: username, PswdHash: string(hashedPassword)}

	user.registerUser()

	c.String(200, "Success")
}

	

func login(c *gin.Context){
	username := c.PostForm("username")
	password := c.PostForm("password")

	failureMessage := "Username or Password is incorrect"
	if usernameExists(username) == false {
		c.JSON(200, gin.H{
			"message": failureMessage,
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(hashedPassword)

	c.JSON(200, gin.H{
		"message": "good",
	})
}