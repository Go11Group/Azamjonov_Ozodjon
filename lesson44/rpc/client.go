package main

import (
	"fmt"
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

	client, err := rpc.DialHTTP("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
	}
	id := GetById{
		id: 1,
	}
	usereponse := ReqUser{}
	//err = client.Call("UserServer.CreateUser", user, &usereponse)
	err = client.Call("UserServer.GetBy", id, &usereponse)
	if err != nil {
		fmt.Println(err)
	}
}
