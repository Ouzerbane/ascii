package main

import (
	"fmt"
	"net/http"

	hand "youmed/handlers"
)

func main() {
	http.HandleFunc("/Download", hand.Download)
	http.HandleFunc("/static/", hand.Static)
	http.HandleFunc("/", hand.HandlerHome) 
	http.HandleFunc("/ascii-art", hand.Handlerascii) 
	fmt.Println("server starting on http://localhost:6969")
	http.ListenAndServe(":6969", nil) 
}