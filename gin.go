package main

import (
	"apigo/wallet"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	router := newServer()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Start the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error: %v", err)
		}
	}()

	// Listen for shutdown signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Server is shutting down...")

	// Give the server a grace period to finish ongoing requests
	const graceTimeout = 5 * time.Second
	ctx, cancel = context.WithTimeout(ctx, graceTimeout)
	defer cancel()

	// Shut down the server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Println("Server gracefully stopped")
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
