package main

import (
	"encoding/json"
	"fmt"
	"io"
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
  "owner": "AnuchitO",
  "balance": 100.0
}

response:
{
	"id": "5f8451e0-3535-4726-b1be-4d152eb3051f",
	"owner": "AnuchitO",
	"balance": 100.0
}
```
*/

type Wallet struct {
	ID      string
	Owner   string
	Balance float64
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Println("name:", name)

	if r.Method == "POST" {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		wt := Wallet{}

		if err := json.Unmarshal(b, &wt); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		res, err := json.Marshal(wt)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
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
	// SELECT * FROM hello WHERE id = {id}
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
