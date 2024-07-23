package models

type UserCreate struct {
	Firstname string
	Lastname  string
	Age       int16
	Role      string
}

type UserGet struct {
	Id        string
	Firstname string
	Lastname  string
	Age       int16
	Role      string
	Time
}

type UserUpdate struct {
	Id        string
	Firstname string
	Lastname  string
	Age       int16
	Role      string
}
