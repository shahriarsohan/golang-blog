package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       uint   `json:"id"`
	Fname    string `json:"first_name"`
	Lname    string `json:"last_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password []byte `json:"-"`
}

func (user *User) SetPassword(password string) {
	hashed_pass, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashed_pass
}
