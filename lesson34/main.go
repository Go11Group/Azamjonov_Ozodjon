package main

import (
	"users_and_produncts/handler"
	"users_and_produncts/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	users := postgres.NewUsersRepo(db)
	products := postgres.NewProdectsRepo(db)

	h := handler.NewHandler(users, products)
	server := handler.CreateServer(h)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	// // Randomly Generate users
	// usersGroup := generator.GenerateUsers(15)
	// for _, user := range usersGroup {
	// 	err := users.CreateUser(user)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// Randomly Generate products
	// ProductsGroup := generator.GenerateProducts(30)
	// for _, product := range ProductsGroup {
	// 	err := products.CreateProduct(product)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
}
