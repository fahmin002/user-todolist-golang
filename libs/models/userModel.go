package models

import "github.com/google/uuid"

type User struct {
	Id            uuid.UUID `json:"id"`
	Username      string    `json:"username" form:"username"`
	Email         string    `json:"email" form:"email"`
	Password_hash string    `json:"password_hash" form:"password_hash"`
	Password_salt string    `json:"password_salt" form:"password_salt"`
	Created_at    string    `json:"created_at"`
	Updated_at    string    `json:"udpdated_at"`
}
