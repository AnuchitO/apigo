package wallet

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Wallet struct {
	ID      int
	Owner   string
	Balance float64
}

var wallets = make(map[int]Wallet)

func convertID(id string) int {
	// TODO: convert id to int
	i, err := strconv.Atoi(id)
	if err != nil {
		return 0
	}
	return i
}

func GetWalletByIDHandler(c *gin.Context) {
	id := convertID(c.Param("id"))
	wt, ok := wallets[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "wallet not found",
		})
		return
	}
	c.JSON(http.StatusOK, wt)
}

func GetBalanceByIDHandler(c *gin.Context) {
	id := convertID(c.Param("id"))
	wt, ok := wallets[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "wallet not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":      wt.ID,
		"balance": wt.Balance,
	})
}

func DepositByIDHandler(c *gin.Context) {
	id := convertID(c.Param("id"))
	wt, ok := wallets[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "wallet not found",
		})
		return
	}

	var payload struct {
		Amount float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	wt.Balance += payload.Amount
	wallets[id] = wt

	c.JSON(http.StatusOK, gin.H{
		"id":      wt.ID,
		"balance": wt.Balance,
	})
}

func WithdrawByIDHandler(c *gin.Context) {
	id := convertID(c.Param("id"))
	wt, ok := wallets[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "wallet not found",
		})
		return
	}

	var payload struct {
		Amount float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if wt.Balance < payload.Amount {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "balance not enough",
		})
		return
	}

	wt.Balance -= payload.Amount
	wallets[id] = wt

	c.JSON(http.StatusOK, gin.H{
		"id":      wt.ID,
		"balance": wt.Balance,
	})
}
