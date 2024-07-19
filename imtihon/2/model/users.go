package model

import "time"

type Users struct {
	UserId    string    `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Birthday  string    `json:"birthday"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt int       `json:"deleted_at"`
}
