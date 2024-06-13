package main

import (
	"github.com/Go11Group/at_lesson/lesson37/gin"
	"github.com/Go11Group/at_lesson/lesson37/psql"
)

func main() {
	db, err := psql.ConnectDb()
	if err != nil {
		panic(err)
	}
	user := psql.CreateUser(db)
	start := gin.ConnectGin(gin.Handler{User: user})
	start.Run()
}
