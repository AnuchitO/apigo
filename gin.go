package main

import (
	"apigo/wallet"
	"log"

	"github.com/gin-gonic/gin"
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

func main() {
	r := newServer()
	r.Run() // listen and serve on 0.0.0.0:8080
}

// Logger handler
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("handler:", c.Request.Method, c.Request.URL)
		c.Next()
		log.Println("after call next handler:", c.Request.Method, c.Request.URL)
	}
}

// wrapper gin.HandlerFunc to use custom logger
// take gin.HandlerFunc as input and return gin.HandlerFunc
func Log(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Auth:")
		// if not pass auth
		// c.Abort()
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		// return

		next(c)
		log.Println("xxxx call next handler:", c.Request.Method, c.Request.URL)
	}
}

// new Server return Gin
func newServer() *gin.Engine {
	r := gin.Default()
	r.Use(Logger())

	r.POST("/wallets", Log(wallet.CreateWalletHandler))
	r.GET("/wallets/:id", wallet.GetWalletByIDHandler)
	r.GET("/wallets/:id/balance", wallet.GetBalanceByIDHandler)
	r.POST("/wallets/:id/deposit", wallet.DepositByIDHandler)
	r.POST("/wallets/:id/withdraw", wallet.WithdrawByIDHandler)

	return r
}
