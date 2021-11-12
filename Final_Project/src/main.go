package main

import (
	"Go/Final_Project/src/controllers/clientcontroller"
	"net/http"
)

func main() {
	http.HandleFunc("/", clientcontroller.Index)
	http.HandleFunc("/client", clientcontroller.Index)
	http.HandleFunc("/client/index", clientcontroller.Index)

	http.ListenAndServe("8080", nil)
}
