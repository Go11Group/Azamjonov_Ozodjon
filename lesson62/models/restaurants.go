package models

type RestaurantCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Address     string `json:"address"`
	ProneNumber string `json:"phone_number"`
}

type RestaurantGet struct {
	Id          string
	Title       string
	Description string
	Address     string
	ProneNumber string
	Time
}

type RestaurantUpdate struct {
	Id          string
	Title       string
	Description string
	Address     string
	ProneNumber string
}
