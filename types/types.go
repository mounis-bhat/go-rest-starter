package types

import (
	"time"
)

type CreateAccountRequest struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Roles     []string `json:"roles"`
}

func NewAccount(firstName string, lastName string, username string, password string, roles []string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Password:  password,
		Roles:     roles,
		CreatedAt: time.Now().UTC(),
	}
}

type ValidateAccountRequest struct {
	ID        int    `json:"account_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Balance   int    `json:"balance"`
	CreatedAt string `json:"created_at"`
	Username  string `json:"username"`
	Roles     string `json:"roles"`
	IAT       int    `json:"iat"`
	EXP       int    `json:"exp"`
}

type Account struct {
	ID        int       `json:"account_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Roles     []string  `json:"roles"`
}
