package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello")
}

func swap(x, y string) (string, string) {
	return y, x
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

func Show(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Show a List of Customer!"))
}

func Server() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/show", Show)
	fmt.Println("Listening on Port :8080 ...")
	http.ListenAndServe(":8080", nil)
}