package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TODO: Transactions
/*
POST /wallets
```json
POST /wallets
payload:
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


GET /wallets/:id
```json
GET /wallets/5f8451e0-3535-4726-b1be-4d152eb3051f
payload: no payload
response:

	{
		"id": "5f8451e0-3535-4726-b1be-4d152eb3051f",
		"owner": "AnuchitO",
		"balance": 100.0
	}

```

GET /wallets/:id/balance
```json
GET /wallets/5f8451e0-3535-4726-b1be-4d152eb3051f/balance
payload: no payload
response:

	{
		"id": "5f8451e0-3535-4726-b1be-4d152eb3051f",
		"balance": 100.0
	}

```

POST /wallets/:id/deposit
```json
POST /wallets/5f8451e0-3535-4726-b1be-4d152eb3051f/deposit
payload:
	{
		"amount": 100.0
	}

response:

	{
		"id": "5f8451e0-3535-4726-b1be-4d152eb3051f",
		"balance": 200.0
	}

```

POST /wallets/:id/withdraw
```json
POST /wallets/5f8451e0-3535-4726-b1be-4d152eb3051f/withdraw
payload:
	{
		"amount": 100.0
	}

response:

	{
		"id": "5f8451e0-3535-4726-b1be-4d152eb3051f",
		"balance": 100.0
	}

```
*/

type Wallet struct {
	ID      string  `json:"id"`
	Owner   string  `json:"owner"`
	Balance float64 `json:"balance"`
}

// Wallet dababase in memory
var wallets = make(map[string]Wallet)

// var wallets = []Wallet{}

func createWalletHandler(c *gin.Context) {
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

func getWalletByIDHandler(c *gin.Context) {
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

func getBalanceByIDHandler(c *gin.Context) {
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

func depositByIDHandler(c *gin.Context) {
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

func withdrawByIDHandler(c *gin.Context) {
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

func main() {
	r := gin.Default()

	r.POST("/wallets", createWalletHandler)
	r.GET("/wallets/:id", getWalletByIDHandler)
	r.GET("/wallets/:id/balance", getBalanceByIDHandler)
	r.POST("/wallets/:id/deposit", depositByIDHandler)
	r.POST("/wallets/:id/withdraw", withdrawByIDHandler)

	r.Run() // listen and serve on 0.0.0.0:8080
}
