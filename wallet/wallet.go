package wallet

import (
	"strconv"
)

type Wallet struct {
	ID      int
	Owner   string
	Balance float64
}

var wallets = make(map[int]Wallet)

func convertID(id string) int {
	i, err := strconv.Atoi(id)
	if err != nil {
		return 0
	}
	return i
}
