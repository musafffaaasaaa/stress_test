package main

import (
	"awesomeProject3/project/handler"
	"net/http"
)

func main() {

	handler.NewHandler()

	http.HandleFunc("/shop/add", func(w http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/shop/get", func(w http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/shop/delete", func(w http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/shop/drop", func(w http.ResponseWriter, r *http.Request) {

	})

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}
