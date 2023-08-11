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
	router := newServer()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// Start listening for shutdown signals
	go func() {
		log.Println("waiting..")
		<-shutdown
		// Close database connections, channels, etc.
		log.Println("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Server shutdown error: %v", err)
		}
		log.Println("Shutting down gracefully...")
		os.Exit(0)
	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error: %v", err)
	}
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
