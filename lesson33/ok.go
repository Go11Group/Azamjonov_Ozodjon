package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()
	//http://localhost:8080/send-url?name=Asadbe&age=15
	mux.HandleFunc("/send-url", sendURL)

	mux.HandleFunc("/send-body", sendBody)
	//http://localhost:8080/send-param?name=Aka&age=5
	mux.HandleFunc("/send-param", sendParam)
	mux.HandleFunc("/get-url", getURL)
	mux.HandleFunc("/get-body", getBody)
	mux.HandleFunc("/get-param", getParam)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
	}
}

// Handler for sending information via URL
func sendURL(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	fmt.Fprintf(w, "Received via URL - Name: %s, Age: %s", name, age)
}

// Handler for sending information via Body
func sendBody(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Received via Body - Data: %s", string(body))
}

// Handler for sending information via URL parameters
func sendParam(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	fmt.Fprintf(w, "Received via URL Params - Name: %s, Age: %s", name, age)
}

// Handler for reading information from URL
func getURL(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	age := query.Get("age")
	if name == "" {
		name = "unknown"
	}
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		ageInt = 0
	}
	fmt.Fprintf(w, "Received from URL - Name: %s, Age: %d", name, ageInt)
}

// Handler for reading information from Body
func getBody(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var body RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Unable to parse JSON request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Received from Body - Name: %s, Age: %d", body.Name, body.Age)
}

// Handler for reading information from URL parameters
func getParam(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	fmt.Fprintf(w, "Received from URL Params - Name: %s, Age: %s", name, age)
}
