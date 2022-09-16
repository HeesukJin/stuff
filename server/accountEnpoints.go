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

	if userExists(username).Username != "" {
		c.String(401, "User already exists.")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		c.String(500, "Something went wrong on our end, please try again.")
		return
	}

	user := User{Username: username, HashedPwd: hashedPassword}
	fmt.Println(user.HashedPwd)

	user.registerUser()

	c.String(200, "Success")
}

	

func login(c *gin.Context){
	username := c.PostForm("username")
	password := c.PostForm("password")

	failureMessage := "Username or Password is incorrect."

	user := userExists(username)

	if user.Username == "" {
		c.String(401, failureMessage)
		return
	}
	
	if bcrypt.CompareHashAndPassword(user.HashedPwd, []byte(password)) != nil {
		c.String(401, failureMessage)
		return
	}

	c.String(200, "Success")
}