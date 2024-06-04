package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Product struct {
	gorm.Model
	FirstName  string
	LastName   string
	Email      string
	Password   string
	Age        int
	Field      string
	Gender     string
	IsEmployee bool
}

func connectDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=BEKJONS dbname=ok port=5432 sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

//func createProduct(db *gorm.DB, product Product) {
//	result := db.Create(&product)
//	if result.Error != nil {
//		fmt.Println("Error creating product:", result.Error)
//	} else {
//		fmt.Println("Product created:", product)
//	}
//}

func getProductByID(db *gorm.DB, id uint) (Product, error) {
	var product Product
	result := db.First(&product, id)
	if result.Error != nil {
		return Product{}, result.Error
	}
	return product, nil
}

func updateProduct(db *gorm.DB, id uint, updatedData map[string]interface{}) {
	result := db.Model(&Product{}).Where("id = ?", id).Updates(updatedData)
	if result.Error != nil {
		fmt.Println("Error updating product:", result.Error)
	} else {
		fmt.Println("Product updated")
	}
}

func deleteProduct(db *gorm.DB, id uint) {
	result := db.Delete(&Product{}, id)
	if result.Error != nil {
		fmt.Println("Error deleting product:", result.Error)
	} else {
		fmt.Println("Product deleted")
	}
}

func main() {
	db, err := connectDB()
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	//firstNames := []string{"John", "Jane", "Alice", "Bob", "Charlie", "David", "Eva", "Frank", "Grace", "Henry", "Ivy", "Jack", "Karen", "Larry", "Mona", "Nate", "Olivia", "Paul", "Quincy", "Rachel", "Steve", "Tina", "Uma", "Vince", "Wendy"}
	//lastNames := []string{"Smith", "Johnson", "Williams", "Jones", "Brown", "Davis", "Miller", "Wilson", "Moore", "Taylor", "Anderson", "Thomas", "Jackson", "White", "Harris", "Martin", "Thompson", "Garcia", "Martinez", "Robinson", "Clark", "Rodriguez", "Lewis", "Lee", "Walker"}
	//fields := []string{"Engineering", "Marketing", "Finance", "HR", "Sales"}
	//genders := []string{"Male", "Female"}
	//
	//for i := 0; i < 25; i++ {
	//	product := Product{
	//		FirstName:  firstNames[i],
	//		LastName:   lastNames[i],
	//		Email:      fmt.Sprintf("%s.%s@example.com", firstNames[i], lastNames[i]),
	//		Password:   "password123",
	//		Age:        20 + i,
	//		Field:      fields[i%len(fields)],
	//		Gender:     genders[i%len(genders)],
	//		IsEmployee: i%2 == 0,
	//	}
	//	createProduct(db, product)
	//}

	// Read products
	var retrievedProducts []Product
	result := db.Find(&retrievedProducts)
	if result.Error != nil {
		panic("failed to retrieve data from database")
	}
	for _, product := range retrievedProducts {
		fmt.Printf("ID: %d, FirstName: %s, LastName: %s, Email: %s, Password: %s, Age: %d, Field: %s, Gender: %s, IsEmployee: %t\n",
			product.ID, product.FirstName, product.LastName, product.Email, product.Password, product.Age, product.Field, product.Gender, product.IsEmployee)
	}

	// Update a product
	updateData := map[string]interface{}{
		"LastName": "Smith",
		"Age":      31,
	}
	updateProduct(db, 1, updateData)

	// Read the updated product
	updatedProduct, err := getProductByID(db, 1)
	if err != nil {
		fmt.Println("Error retrieving updated product:", err)
	} else {
		fmt.Printf("Updated Product: %+v\n", updatedProduct)
	}

	// Delete a product
	deleteProduct(db, 1)

	// Verify deletion
	deletedProduct, err := getProductByID(db, 1)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Deleted Product not found:", err)
		} else {
			fmt.Println("Error retrieving deleted product:", err)
		}
	} else {
		fmt.Printf("Deleted Product: %+v\n", deletedProduct)
	}
}
