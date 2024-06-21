package main

import (
	"database/sql"
	"fmt"
	"my_project/api"
	"my_project/strorage/postgres"
)

/*
	bu yerda main function yani run jarayonida shu yerda

sodir boladi chunki  yana postgres ichidagi connectionni shu yerda qilingan
onnection ham
*/
func main() {

	db, err := postgres.ConnectionDb() // postgres connectiondb chaqirisilishi va keyin methodlarga shu yerda connection qilib yuborilishi
	if err != nil {
		panic(err) // agar connection bolmasa panic yani dastur toxtaydi
	}
	defer func(db *sql.DB) {
		err := db.Close() // bu yerda  sql ni close qilib kelasmiz yani defer return dan oldin close qilinadi
		if err != nil {

		}
	}(db)
	fmt.Println("Connect!!!")

	user := postgres.NewUserRepository(db)             // user postgres ichida methodlarni database bilan connection ni ulash
	course := postgres.NewCourseRepository(db)         // user postgres ichida methodlarni database bilan connection ni ulsah
	lesson := postgres.NewLessonRepository(db)         // user postgres ichida methodlarni database bilan connection ni ulsah
	enrollment := postgres.NewEnrollmentRepository(db) // user postgres ichida methodlarni database bilan connection ni ulsah

	router := api.RooterApi(course, enrollment, lesson, user) // bu yerda api larimiz keladi router ga yozilgan api lar

	err = router.Run(":8080") // dastur localhost 8090 da run boladi
	if err != nil {
		panic(err) //error bolsa dastur toxtaydi asosan port band bolsa error beradi
	}

}
