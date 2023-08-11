package main

import (
	"apigo/wallet"
	"log"
	"os"
	"os/signal"
	"syscall"

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
		"id": "9",
		"owner": "AnuchitO",
		"balance": 100.0
	}

```


GET /wallets/:id
```json
GET /wallets/9
payload: no payload
response:

	{
		"id": "9",
		"owner": "AnuchitO",
		"balance": 100.0
	}

```

GET /wallets/:id/balance
```json
GET /wallets/9/balance
payload: no payload
response:

	{
		"id": "9",
		"balance": 100.0
	}

```

POST /wallets/:id/deposit
```json
POST /wallets/9/deposit
payload:
	{
		"amount": 100.0
	}

response:

	{
		"id": "9",
		"balance": 200.0
	}

```

POST /wallets/:id/withdraw
```json
POST /wallets/9/withdraw
payload:
	{
		"amount": 100.0
	}

response:

	{
		"id": "9",
		"balance": 100.0
	}

```

give me a curl to test all of these endpoints

curl -X POST -H "Content-Type: application/json" -d '{"owner": "AnuchitO", "balance": 100.0}' http://localhost:8080/wallets
curl -X GET http://localhost:8080/wallets/9
curl -X GET http://localhost:8080/wallets/9/balance
curl -X POST -H "Content-Type: application/json" -d '{"amount": 100.0}' http://localhost:8080/wallets/9/deposit
curl -X POST -H "Content-Type: application/json" -d '{"amount": 100.0}' http://localhost:8080/wallets/9/withdraw


*/

func main() {
	r := newServer()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// Start listening for shutdown signals
	go func() {
		log.Println("waiting..")
		<-shutdown
		// Close database connections, channels, etc.

		log.Println("Shutting down gracefully...")
		os.Exit(0)
	}()

	r.Run() // listen and serve on 0.0.0.0:8080
}

// new Server return Gin
func newServer() *gin.Engine {
	r := gin.Default()

	r.POST("/wallets", wallet.CreateWalletHandler)
	r.GET("/wallets/:id", wallet.GetWalletByIDHandler)
	r.GET("/wallets/:id/balance", wallet.GetBalanceByIDHandler)
	r.POST("/wallets/:id/deposit", wallet.DepositByIDHandler)
	r.POST("/wallets/:id/withdraw", wallet.WithdrawByIDHandler)

	return r
}
