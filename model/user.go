package model

import (
	"errors"
	"strings"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

var userList = []User{
	User{
		Username: "20th",
		Password: "nclab",
	},
}

func IsUserValid(username, password string) bool {
	for _, u := range userList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func IsUsernameAvailable(username string) bool {
	for _, u := range userList {
		if u.Username == username {
			return false
		}
	}
	return true
}

func RegisterNewUser(username, password string) (*User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("The password can't be empty")
	} else if !IsUsernameAvailable(username) {
		return nil, errors.New("The username isn't available")
	}

	u := User{
		Username: username,
		Password: password,
	}

	userList = append(userList, u)

	return &u, nil
}
