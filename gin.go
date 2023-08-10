package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET http://localhost:8080/hello?name=nong
// GET http://localhost:8080/hello/anuchito?name=nong

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
