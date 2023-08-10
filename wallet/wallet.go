package wallet

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Wallet struct {
	ID      string  `json:"id"`
	Owner   string  `json:"owner"`
	Balance float64 `json:"balance"`
}

var wallets = make(map[string]Wallet)

func CreateWalletHandler(c *gin.Context) {
	log.Println("handler:", c.Request.Method, c.Request.URL)
	var wt Wallet
	if err := c.ShouldBindJSON(&wt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	wt.ID = uuid.New().String()
	wallets[wt.ID] = wt
	c.JSON(http.StatusCreated, wt)
}

func GetWalletByIDHandler(c *gin.Context) {
	log.Println("handler:", c.Request.Method, c.Request.URL)
	id := c.Param("id")
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
	log.Println("handler:", c.Request.Method, c.Request.URL)
	id := c.Param("id")
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
	log.Println("handler:", c.Request.Method, c.Request.URL)
	id := c.Param("id")
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
	log.Println("handler:", c.Request.Method, c.Request.URL)
	id := c.Param("id")
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
