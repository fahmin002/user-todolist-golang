package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid"`
	Username     string
	Email        *string
	Hashpassword string
	Hashsalt     string
}
