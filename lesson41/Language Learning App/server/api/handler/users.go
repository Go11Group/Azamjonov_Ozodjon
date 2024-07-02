package handler

import (
	"encoding/json"
	"fmt"
	"language_learning_app/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create
func (h *Handler) CreateUser(c *gin.Context) {
	newUser := model.User{}

	// take json from request body to new struct
	err := json.NewDecoder(c.Request.Body).Decode(&newUser) //c.BindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting json from URL body: %s\n", err),
		})
		log.Printf("Error with getting json from URL body: %s\n", err)
		return
	}

	// use Users struct's method to create
	err = h.Users.CreateUser(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with create new user into database: %s", err),
		})
		log.Printf("Error with create new user into database: %s\n", err)
		return
	}

	// send message to response 'there is no error'
	c.JSON(http.StatusOK, "Created new User successfully")
}

// Read all
func (h *Handler) GetUsers(c *gin.Context) {
	filter := model.UserFilter{}

	// get filter parametres from URL query
	name, hasName := c.GetQuery("name")
	if hasName {
		filter.Name = &name
	}
	email, hasEmail := c.GetQuery("email")
	if hasEmail {
		filter.Email = &email
	}
	ageFrom, hasAgeFrom := c.GetQuery("age_from")
	if hasAgeFrom {
		filter.AgeFrom = &ageFrom
	}
	ageTo, hasAgeTo := c.GetQuery("age_to")
	if hasAgeTo {
		filter.AgeTo = &ageTo
	}
	limit, hasLimit := c.GetQuery("limit")
	if hasLimit {
		filter.Limit = &limit
	}
	offset, hasOffset := c.GetQuery("offset")
	if hasOffset {
		filter.Offset = &offset
	}

	// use users struct's mothod to get data from database
	users, err := h.Users.GetUsers(&filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with getting users from database: %s", err),
		})
		log.Printf("Error with getting users from database: %s\n", err)
		return
	}

	// send data in response body
	c.JSON(http.StatusOK, users)
}

// Read by id
func (h *Handler) GetUserById(c *gin.Context) {
	// get id from URL params
	id, hasId := c.Params.Get("id")
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error: id not found in URL params",
		})
		log.Printf("Error: id not found in URL params")
		return
	}

	// use users struct's method to get data by id
	user, err := h.Users.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with getting user by id from database: %s", err),
		})
		log.Printf("Error with getting user by id from database: %s\n", err)
		return
	}

	// send data in response body
	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetUsersCourses(c *gin.Context) {
	userId, hasId := c.Params.Get("id")
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error: id not found in URL params",
		})
		log.Printf("Error: id not found in URL params")
		return
	}
	courses, err := h.Users.GetUsersCourses(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with getting user's courses from database: %s", err),
		})
		log.Printf("Error with getting user's courses from database: %s\n", err)
		return
	}
	c.JSON(http.StatusOK, courses)
}

// Update
func (h *Handler) UpdateUser(c *gin.Context) {
	user := model.User{}

	// get json from request body
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting json from Request body: %s", err),
		})
		log.Printf("Error with getting json from Request body: %s\n", err)
		return
	}

	// get id from URL params
	id, hasId := c.Params.Get("id")
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error: id not found in URL params",
		})
		log.Printf("Error: id not found in URL params")
		return
	}

	// put id to user's userId field
	user.UserId = id

	// use courses struct's method to update the data
	err = h.Users.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with updating user in database: %s", err),
		})
		log.Printf("Error with updating user in database: %s\n", err)
		return
	}

	// send message to response 'there is no error'
	c.JSON(http.StatusOK, "User was updated successfully")
}

// Delete
func (h *Handler) DeleteUser(c *gin.Context) {
	// get id from URL params
	id, hasId := c.Params.Get("id")
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error: id not found in URL params",
		})
		log.Printf("Error: id not found in URL params")
		return
	}

	// use users struct's method to delete the data
	err := h.Users.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with deleting user in database: %s", err),
		})
		log.Printf("Error with deleting user in database: %s\n", err)
		return
	}

	// send message to response 'there is no error'
	c.JSON(http.StatusOK, "User was deleted successfully")
}
