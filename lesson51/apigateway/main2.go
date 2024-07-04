package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
)

type User struct {
	id   int
	name string
	age  int
	jwt.StandardClaims
}

func main() {
	router := gin.Default()
	router.Use(Middleware)
	router.GET("/", getUser)
	router.Run(":8081")
}
func Middleware(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	if tokenStr == "" {
		c.JSON(http.StatusBadRequest, "token not found")

	}
	user := &User{}

	token, err := jwt.ParseWithClaims(tokenStr, user, func(token *jwt.Token) (interface{}, error) {
		return []byte("ok"), nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}
func getUser(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world")
}
