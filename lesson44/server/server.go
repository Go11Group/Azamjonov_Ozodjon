package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type ReqUser struct {
	Id   int
	Name string
}
type ResUser struct {
	Id   int
	Name string
}
type GetById struct {
	id int
}
type UserServer struct{}

var users []ResUser

func main() {
	userService := new(UserServer)
	err := rpc.Register(userService)
	if err != nil {
		fmt.Println("+++", err)

	}
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	panic(http.Serve(listener, nil))
}
func CreateUser(reqUser *ReqUser, resUser *ResUser) {
	users = append(users, ResUser{
		Id:   reqUser.Id,
		Name: reqUser.Name,
	})
	fmt.Println("users", users)
}

func GetBy(rq *GetById, resUser *ResUser) {
	fmt.Println(users[rq.id])
}
