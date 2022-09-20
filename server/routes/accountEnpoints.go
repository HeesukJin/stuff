package routes

import (
	_ "fmt"
	"net/http"
	"tradeout-server/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-contrib/sessions"
)

const userkey = "ID"
var secret = []byte("secret")


func RegisterAccount(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	usernameErr := models.ValidateUsername(username)
	if usernameErr != nil {
		c.String(401, usernameErr.Error())
		return
	}

	passwordErr := models.ValididatePassword(password)
	if passwordErr != nil {
		c.String(401, passwordErr.Error())
		return
	}

	if models.UserExists(username).Username != "" {
		c.String(401, "User already exists.")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		c.String(500, "Something went wrong on our end, please try again.")
		return
	}

	user := models.User{Username: username, HashedPwd: hashedPassword}
	user.RegisterUser()

	c.String(200, "Success")
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	failureMessage := "Username or Password is incorrect."

	user := models.UserExists(username)

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



// Thanks to otraore for the code example
// https://gist.github.com/otraore/4b3120aa70e1c1aa33ba78e886bb54f3

// AuthRequired is a simple middleware to check the session
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete(userkey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func Me(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}
