package generator

import (
	"math/rand"
	"time"
	"users_and_produncts/model"
)

var productNames = []string{
	"Laptop", "Smartphone", "Headphones", "Monitor", "Keyboard", "Mouse",
	"Printer", "Camera", "Speaker", "Router", "Tablet", "Smartwatch",
	"External Hard Drive", "USB Flash Drive", "Gaming Console", "Television",
	"Projector", "Fitness Tracker", "Drone", "Home Assistant",
}

var productDescriptions = []string{
	"High performance and reliability.",
	"Latest model with advanced features.",
	"Compact design with powerful performance.",
	"User-friendly interface and sleek design.",
	"High-quality build and excellent performance.",
	"Affordable price with premium features.",
	"Durable and long-lasting.",
	"Energy-efficient and environmentally friendly.",
	"Compact size with powerful performance.",
	"Versatile and multifunctional.",
	"Easy to use and highly efficient.",
	"Stylish design with cutting-edge technology.",
}

func randomPrice() float64 {
	return float64(rand.Intn(9000)+1000) / 100 // Random price between $10.00 and $100.00
}

func randomStockQuantity() int {
	return rand.Intn(100) + 1 // Random stock quantity between 1 and 100
}

func GenerateProducts(count int) []model.Product {
	rand.Seed(time.Now().UnixNano())

	products := make([]model.Product, count)
	for i := 0; i < count; i++ {
		name := productNames[rand.Intn(len(productNames))]
		description := productDescriptions[rand.Intn(len(productDescriptions))]
		price := randomPrice()
		stockQuantity := randomStockQuantity()

		products[i] = model.Product{
			Name:           name,
			Description:    description,
			Price:          price,
			Stock_quantity: stockQuantity,
		}
	}
	return products
}
