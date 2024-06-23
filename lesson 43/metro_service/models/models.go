package models

import (
	"time"
)

type Card struct {
	Id        string     `json:"id"`
	Number    string     `json:"number"`
	UserId    string     `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type CreateCard struct {
	Number string `json:"number" binding:"required"`
	UserId string `json:"user_id" binding:"required"`
}

type Station struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type CreateStation struct {
	Name string `json:"name" binding:"required"`
}

type Terminal struct {
	Id        string     `json:"id"`
	StationId string     `json:"station_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type CreateTerminal struct {
	StationId string `json:"station_id" binding:"required"`
}

type Transaction struct {
	Id              string     `json:"id"`
	CardId          string     `json:"card_id"`
	Amount          float64    `json:"amount"`
	TerminalId      *string    `json:"terminal_id"`
	TransactionType string     `json:"transaction_type"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}

type CreateTransaction struct {
	CardId          string  `json:"card_id" binding:"required"`
	Amount          float64 `json:"amount" binding:"required"`
	TerminalId      *string `json:"terminal_id"`
	TransactionType string  `json:"transaction_type" binding:"required"`
}

type User struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	Phone     string     `json:"phone"`
	Age       int        `json:"age"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type CreateUser struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Age   int    `json:"age" binding:"required"`
}
