package model

import "database/sql"

type Product struct {
	Time
	Id             int
	Name           string
	Description    sql.NullString
	Price          float64
	Stock_quantity int
}

type FilterProduct struct {
	Id             *int
	Name           *string
	Description    *string
	Price          *float64
	Stock_quantity *int
}
