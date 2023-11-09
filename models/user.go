package models

import "fmt"

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Posts    []Post `json:"posts"`
}

func (user *User) Create() error {
	fmt.Println("Creating a user", "user email:", user.Email, "user password:", user.Password, "user name:", user.Name)
	return nil
}
