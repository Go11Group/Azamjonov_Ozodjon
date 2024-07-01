package handler

import (
	"context"
	"fmt"
	pb "homeworks/homework_47/genproto/transport"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *HandlerRepo) GetBusSchedule(c *gin.Context) {
	a := c.Param("id")
	id, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": "error in Atoi"})
	}

	req := pb.BusNumber{Bus: int32(id)}
	ctx := context.Background()
	res, err := h.Transport.GetBusSchedule(ctx, &req)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": "error in get bus schedule"})
	}
	c.JSON(http.StatusOK, res)
}

func (h *HandlerRepo) TrucBusLocation(c *gin.Context) {
	req := pb.Location{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": err.Error()})
		return
	}
	ctx := context.Background()
	res, err := h.Transport.TrackBusLocation(ctx, &req)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
