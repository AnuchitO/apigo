package main

import "fmt"

/*
{"id": "wallet_AnuchitO","owner": "AnuchitO","balance": 100.0}
*/

type wallet struct {
	ID      string  `json:"id"`
	Owner   string  `json:"account"`
	Balance float64 `json:"balance"`
}

func main() {
	var wt wallet
	fmt.Printf("%#v\n", wt)
	fmt.Printf("type: %T\n", wt)
	i := 45
	p := &i
	fmt.Println(*p)
	var wt2 = &wallet{}
	fmt.Printf("type: %T\n", wt2)
	var wt3 = new(wallet)
	fmt.Printf("type: %T\n", wt3)
}
