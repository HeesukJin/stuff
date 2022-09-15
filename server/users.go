package main

import(
	"unicode"
	"fmt"
	"database/sql"

	passwordvalidator "github.com/wagslane/go-password-validator"
	"errors"

)

func (u *User) getUserByUsername() error {
	stmt := "SELECT * FROM users WHERE username = ?"
	row := db.QueryRow(stmt, u.Username)
	err := row.Scan(&u.ID, &u.Username, &u.Email, &u.pswdHash, &u.CreatedAt, &u.Active, &u.verHash, &u.timeout)
	if err != nil {
		fmt.Println("getUser() error selecting User, err:", err)
		return err
	}
	return nil
}

func(u *User) validPassword() error {
	// if the password has enough entropy, err is nil
	// otherwise, a formatted error message is provided explaining
	// how to increase the strength of the password
	// (safe to show to the client)
	err := passwordvalidator.Validate(u.password, 60)
	return err
	}


// validateUsername checks username only has alphanumeric characters
// and if sufficient length, errors are safe to share with user
func (u *User) validateUsername() error {
	// check username for only alphaNumeric characters
	for _, char := range u.Username {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			return errors.New("only alphanumeric characters allowed for username")
		}
	}
	// check username length
	if 5 <= len(u.Username) && len(u.Username) <= 50 {
		return nil
	}
	return errors.New("username length must be greater than 4 and less than 51 characters")
}

func (u *User) UsernameExists() (exists bool) {
	exists = true
	stmt := "SELECT id FROM users WHERE username = ?"
	row := db.QueryRow(stmt, u.Username)
	var uID string
	err := row.Scan(&uID)
	if err == sql.ErrNoRows {
		return false
	}
	return exists
}
