package main

import (
	"encoding/json"
	"log"
)

/*
{"id": "wallet_AnuchitO","owner": "AnuchitO","balance": 100.0}
*/

type wallet struct {
	ID      string  `json:"id"`
	Owner   string  `json:"account"`
	Balance float64 `json:"balance"`
}

func main() {
	b := `{"ID": "wallet_AnuchitO","account": "AnuchitO","balance": 100.0}`
	wt := wallet{}
	err := json.Unmarshal([]byte(b), &wt)
	if err != nil {
		log.Println("err:", err)
		return
	}
	log.Printf("wt: %#v\n", wt)

	res, _ := json.Marshal(wt)
	log.Printf("Marshal: %#v\n", string(res))
}
