package main

import(
	"unicode"
	"fmt"
	"database/sql"

	passwordvalidator "github.com/wagslane/go-password-validator"
	"errors"

)

type User struct {
	ID        string
	Username  string
	Email     string
	PswdHash  string
	CreatedAt string
	Active    string
	verHash   string
	timeout   string
}

func (u *User) registerUser() error {
	stmt := "INSERT INTO users VALUES(?, ?)"
	res, err := MySQLClient.Exec(stmt, u.Username, u.PswdHash)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
	return nil
}


func (u *User) getUserByUsername() error {
	stmt := "SELECT * FROM users WHERE username = ?"
	row := MySQLClient.QueryRow(stmt, u.Username)
	err := row.Scan(&u.ID, &u.Username, &u.Email, &u.PswdHash, &u.CreatedAt, &u.Active, &u.verHash, &u.timeout)
	if err != nil {
		fmt.Println("getUser() error selecting User, err:", err)
		return err
	}

	return nil
}

func valididatePassword(password string) error {
	// if the password has enough entropy, err is nil
	// otherwise, a formatted error message is provided explaining
	// how to increase the strength of the password
	// (safe to show to the client)
	err := passwordvalidator.Validate(password, 60)
	return err
}


// validateUsername checks username only has alphanumeric characters
// and if sufficient length, errors are safe to share with user
func validateUsername(username string) error {
	// check username for only alphaNumeric characters
	for _, char := range username {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			return errors.New("only alphanumeric characters allowed for username")
		}
	}
	// check username length
	if 5 <= len(username) && len(username) <= 50 {
		return nil
	}
	return errors.New("username length must be greater than 4 and less than 51 characters")
}

func usernameExists(username string) (exists bool) {
	exists = true
	stmt := "SELECT username FROM users WHERE username = ?"
	row := MySQLClient.QueryRow(stmt, username)
	var uID string

	if row.Scan(&uID) == sql.ErrNoRows {
		return false
	}
	return exists
}
