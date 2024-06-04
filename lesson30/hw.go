package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "BEKJONS"
	dbname   = "ok"
)

var db *sql.DB

func initDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database")
}

func createUser(username, email, password string) {
	_, err := db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", username, email, password)
	if err != nil {
		log.Printf("Error creating user: %v\n", err)
	} else {
		fmt.Println("User created successfully")
	}
}

func createProduct(name, description string, price float64, stockQuantity int) {
	_, err := db.Exec("INSERT INTO products (name, description, price, stock_quantity) VALUES ($1, $2, $3, $4)", name, description, price, stockQuantity)
	if err != nil {
		log.Printf("Error creating product: %v\n", err)
	} else {
		fmt.Println("Product created successfully")
	}
}

func getUser(userID int) {
	var id int
	var username, email, password string
	err := db.QueryRow("SELECT id, username, email, password FROM users WHERE id = $1", userID).Scan(&id, &username, &email, &password)
	if err != nil {
		log.Printf("Error reading user: %v\n", err)
	} else {
		fmt.Printf("User: ID=%d, Username=%s, Email=%s, Password=%s\n", id, username, email, password)
	}
}

func getProduct(productID int) {
	var id int
	var name, description string
	var price float64
	var stockQuantity int
	err := db.QueryRow("SELECT id, name, description, price, stock_quantity FROM products WHERE id = $1", productID).Scan(&id, &name, &description, &price, &stockQuantity)
	if err != nil {
		log.Printf("Error reading product: %v\n", err)
	} else {
		fmt.Printf("Product: ID=%d, Name=%s, Description=%s, Price=%.2f, StockQuantity=%d\n", id, name, description, price, stockQuantity)
	}
}

func updateUser(userID int, username, email, password string) {
	_, err := db.Exec("UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4", username, email, password, userID)
	if err != nil {
		log.Printf("Error updating user: %v\n", err)
	} else {
		fmt.Println("User updated successfully")
	}
}

func updateProduct(productID int, name, description string, price float64, stockQuantity int) {
	_, err := db.Exec("UPDATE products SET name = $1, description = $2, price = $3, stock_quantity = $4 WHERE id = $5", name, description, price, stockQuantity, productID)
	if err != nil {
		log.Printf("Error updating product: %v\n", err)
	} else {
		fmt.Println("Product updated successfully")
	}
}

func deleteUser(userID int) {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		log.Printf("Error deleting user: %v\n", err)
	} else {
		fmt.Println("User deleted successfully")
	}
}

func deleteProduct(productID int) {
	_, err := db.Exec("DELETE FROM products WHERE id = $1", productID)
	if err != nil {
		log.Printf("Error deleting product: %v\n", err)
	} else {
		fmt.Println("Product deleted successfully")
	}
}

func transactionExample(users []map[string]string, products []map[string]interface{}) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Error starting transaction: %v\n", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			log.Printf("Error in transaction: %v\n", err)
			tx.Rollback()
		} else {
			err = tx.Commit()
			if err != nil {
				log.Printf("Error committing transaction: %v\n", err)
			}
		}
	}()

	for _, user := range users {
		_, err = tx.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user["username"], user["email"], user["password"])
		if err != nil {
			return
		}
	}

	for _, product := range products {
		_, err = tx.Exec("INSERT INTO products (name, description, price, stock_quantity) VALUES ($1, $2, $3, $4)", product["name"], product["description"], product["price"], product["stock_quantity"])
		if err != nil {
			return
		}
	}
}

func main() {
	initDB()

	users := []map[string]string{
		{"username": "user1", "email": "user1@example.com", "password": "password1"},
		{"username": "user2", "email": "user2@example.com", "password": "password2"},
	}

	products := []map[string]interface{}{
		{"name": "product1", "description": "description1", "price": 10.00, "stock_quantity": 100},
		{"name": "product2", "description": "description2", "price": 20.00, "stock_quantity": 200},
	}

	transactionExample(users, products)

	getUser(1)
	getProduct(1)
	updateUser(1, "updated_user", "updated_email@example.com", "updated_password")
	updateProduct(1, "updated_product", "updated_description", 15.00, 150)
	deleteUser(2)
	deleteProduct(2)
}
