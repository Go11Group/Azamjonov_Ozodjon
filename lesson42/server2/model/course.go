package model

import (
	"time"
)

// Course struct  elen qilish   bu
// struct biz ga table dan malumot o'qib response
// ruborish uchun yoki  client dan request olib birinchi  structg parse qilib keyin json
type Course struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   int       `json:"deleted_at"`
}
