package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"my_project/api/handlers"
	"net/http"
)

func HandleHTTPResponse(gn *gin.Context, response *http.Response, err error) {
	if err != nil {
		handlers.ErrorResponse(gn, err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		var result map[string]interface{}
		if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
			handlers.ErrorResponse(gn, err)
			return
		}
		gn.JSON(response.StatusCode, result)
	} else {
		var errorResult map[string]interface{}
		if err := json.NewDecoder(response.Body).Decode(&errorResult); err != nil {
			handlers.ErrorResponse(gn, fmt.Errorf(" status code: %d", response.StatusCode))
			return
		}
		gn.JSON(response.StatusCode, errorResult)
	}
}
