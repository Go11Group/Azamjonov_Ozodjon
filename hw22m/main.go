package main

import (
	"fmt"
	"mod_my/task"
)

func main() {
	
	fmt.Println("add =", task.Add(10, 8))
	fmt.Println("min =", task.Min(10, 8))	
}