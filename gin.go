package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: Transactions
/*
POST /wallets
```json
POST /wallets

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

response:

	{
		"id": "5f8451e0-3535-4726-b1be-4d152eb3051f",
		"balance": 100.0
	}

```

POST /wallets/:id/deposit
```json
POST /wallets/5f8451e0-3535-4726-b1be-4d152eb3051f/deposit

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

// Wallet dababase in memory
// var wallets = make(map[string]Wallet)
// var wallets = []Wallet{}

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		name := c.Query("name")
		c.JSON(200, gin.H{
			"message": name,
		})
	})

	r.GET("/hello/:id", func(c *gin.Context) {
		c.Request.URL.Query().Get("name")
		name := c.Query("name")

		id := c.Param("id")

		c.JSON(200, gin.H{
			"id":   id,
			"name": name,
		})
	})
	/*
		```json
		POST /wallets
		{
			"id": "wallet_AnuchitO",
			"owner": "AnuchitO",
			"balance": 100.0
		}
		```
	*/

	r.POST("/wallets", func(c *gin.Context) {
		var wt Wallet
		if err := c.ShouldBindJSON(&wt); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, wt)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
