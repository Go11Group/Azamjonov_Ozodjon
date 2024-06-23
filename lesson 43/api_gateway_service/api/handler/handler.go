package handler

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) Client(ctx *gin.Context) {
	method := ctx.Request.Method
	url := ctx.Request.URL.Path

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error reading body:", err.Error())
		return
	}

	client := http.Client{}
	req, err := http.NewRequest(method, "http://localhost:8080"+url, bytes.NewBuffer(body))
	fmt.Println("http://localhost:8080" + url)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error creating request:", err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error doing request:", err.Error())
		return
	}

	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error reading response:", err.Error())
		return
	}

	// Copy status code from the response
	ctx.JSON(res.StatusCode, gin.H{"data": string(respBody)})
}
