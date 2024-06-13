package model

type Users struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Birthday  string  `json:"birthday"`
	Password  string  `json:"password"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}
