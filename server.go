package main

import (
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func main() {
	log.Println("Starting server on port 8080")

	http.HandleFunc("/hello", helloHandler)

	log.Println("Goodbye!")
}
