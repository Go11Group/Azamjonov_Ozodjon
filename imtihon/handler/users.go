package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imtihon/model"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// UserGetById retrieves a user by ID.
func (h *Handler) UserGetById(c *gin.Context) {
	id := c.Param("id")
	user, err := h.Users.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UserGet retrieves users based on query parameters.
func (h *Handler) UserGet(c *gin.Context) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset = ""
	)

	query := `select * from users where deleted_at=0`

	// Apply filters based on query parameters
	if len(c.Query("name")) > 0 {
		params["name"] = c.Query("name")
		query += " and name = :name"
	}

	if len(c.Query("email")) > 0 {
		params["email"] = c.Query("email")
		query += " and email = :email "
	}

	if len(c.Query("birthday")) > 0 {
		params["birthday"] = c.Query("birthday")
		query += " and birthday = :birthday "
	}

	// Apply limit and offset if specified
	if len(c.Query("limit")) > 0 {
		params["limit"] = c.Query("limit")
		limit = ` LIMIT :limit`
	}

	if len(c.Query("offset")) > 0 {
		params["offset"] = c.Query("offset")
		offset = ` OFFSET :offset`
	}

	// Construct final query with optional limit and offset
	query = query + limit + offset

	// Replace named query parameters with positional parameters
	query, arr = ReplaceQueryParamsU(query, params)

	// Execute query to get users from database
	users, err := h.Users.Get(query, arr)
	if err != nil {
		fmt.Println(err) // Print error to console
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR IN GET": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// ReplaceQueryParamsU converts named query parameters to positional parameters.
func ReplaceQueryParamsU(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)
	for k, v := range params {
		if k != "" && strings.Contains(namedQuery, ":"+k) {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))
			args = append(args, v)
			i++
		}
	}
	return namedQuery, args
}

// UserCreate creates a new user.
func (h *Handler) UserCreate(c *gin.Context) {
	var user model.Users

	// Bind JSON request body to user model
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set CreatedAt and UpdatedAt timestamps
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Call UserCreate method in repository
	if err := h.Users.UserCreate(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// UserUpdate updates an existing user.
func (h *Handler) UserUpdate(c *gin.Context) {
	id := c.Param("id")
	var user model.Users

	// Bind JSON request body to user model
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set UserId and UpdatedAt timestamp
	user.UserId = id
	user.UpdatedAt = time.Now()

	// Call UserUpdate method in repository
	if err := h.Users.UserUpdate(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UserDelete deletes a user by ID.
func (h *Handler) UserDelete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Users.UserDelete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// UserSearch searches users by name and/or email.
func (h *Handler) UserSearch(c *gin.Context) {
	name := c.Query("name")
	email := c.Query("email")

	// Call SearchUsers method in repository
	users, err := h.Users.SearchUsers(name, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": users})
}
