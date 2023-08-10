package main

import (
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("helloHandler")
	w.Write([]byte("Hello, world!"))
}

func main() {
	log.Println("Starting server on port 8080")

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
