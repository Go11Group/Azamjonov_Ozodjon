package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// status lar ok,badRequest,InternalSeverError
// InternalSeverError
func ErrorResponse(gn *gin.Context, err error) {

	gn.JSON(http.StatusInternalServerError, gin.H{
		"message": fmt.Sprintf("%s", err),
		"status":  http.StatusInternalServerError,
		"time":    time.Now(),
	})
}

// ok
func Ok(gn *gin.Context) {
	gn.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s", "Success"),
		"status":  http.StatusOK,
		"time":    time.Now(),
	})
}

// badRequest
func BadRequest(gn *gin.Context, err error) {
	gn.JSON(http.StatusBadRequest, gin.H{
		"message": fmt.Sprintf("%s", err),
		"status":  http.StatusBadRequest,
		"time":    time.Now(),
	})
}
