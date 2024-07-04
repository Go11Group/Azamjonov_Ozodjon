package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func main() {
	router := gin.Default()
	router.POST("/user", userHandler)
	router.Run(":8080")
}

func userHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Name != "admin" || user.Password != "admin" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Name
	claims["id"] = user.ID
	claims["age"] = user.Age
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["iat"] = time.Now().Unix()

	newToken, err := token.SignedString([]byte("ok"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": newToken})
}
