package main

import "client/api"

func main() {
	panic(api.Routes().ListenAndServe())
}
