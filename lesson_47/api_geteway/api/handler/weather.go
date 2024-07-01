package handler

import (
	"context"
	"fmt"
	pb "homeworks/homework_47/genproto/weather"
	"net/http"

	"github.com/gin-gonic/gin"
)

func(h *HandlerRepo) GetCurrentWeather(c *gin.Context){
	req := pb.Void{}
	ctx := context.Background()
	res, err := h.Weather.GetCurrentWeather(ctx, &req)
	if err != nil{
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": "error in get current weather"})
	}
	c.JSON(http.StatusOK, res)
}