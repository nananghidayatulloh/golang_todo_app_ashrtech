package model

import (
	"time"
)

type User struct {
	Base
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
	Todos           []Todo `json:"todos"`
}

type Todo struct {
	Base
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      uint   `json:"userid"`
	Start       string `json:"start"`
	End         string `json:"end"`
	Status      string `json:"status"`
}

type Base struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
