package main

import (
	"my_project/api"
)

func main() {

	router := api.RooterApi() // bu yerda api larimiz keladi router ga yozilgan api lar

	err := router.Run(":8090") // dastur localhost 8090 da run boladi
	if err != nil {
		panic(err) //error bolsa dastur toxtaydi asosan port band bolsa error beradi
	}

}
