package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
POST /wallets
GET /wallets/:id/balance
POST /wallets/:id/deposit
POST /wallets/:id/withdraw

```json
POST /wallets
{
  "id": "wallet_AnuchitO",
  "owner": "AnuchitO",
  "balance": 100.0,
}
```
*/

type Wallet struct {
	ID      string
	Owner   string
	Balance float64
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// get query param from http://localhost:8080/users/:id?name=nong
	// get query param from http://localhost:8080/users/:id
	name := r.URL.Query().Get("name")
	fmt.Println("name:", name)

	if r.Method == "POST" {
		// create wallet
		w.Write([]byte("created"))
		return
	}

	if r.Method == "GET" {
		w.Write([]byte("Hello, world!"))
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func main() {
	log.Println("Starting server on port 8080")

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/hello/{id}", helloHandler)
	// SELECT * FROM hello WHERE id = {id}
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
