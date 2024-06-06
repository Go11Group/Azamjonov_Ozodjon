package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", hello)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {

	fmt.Println("URL:", r.URL)
	fmt.Println("Host:", r.Host)
	fmt.Println("Method:", r.Method)

	_, err := w.Write([]byte("IN HELLOOOO"))
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}
