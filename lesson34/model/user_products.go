package model

type User_product struct {
	Time
	Id         int
	User_id    int
	Product_id int
}

type FilterUserProduct struct {
	Id         *int
	User_id    *int
	Product_id *int
}
