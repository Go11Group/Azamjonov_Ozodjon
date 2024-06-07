package model

type User struct {
	Time
	Id       int
	Username string
	Email    string
	Password string
}

type Filter struct {
	Id       *int
	Username *string
}
