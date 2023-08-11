package wallet

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateWalletHandler(c *gin.Context) {
	log.Println("Create:")
	var wt Wallet
	if err := c.ShouldBindJSON(&wt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	wt.ID = 0
	wallets[wt.ID] = wt
	c.JSON(http.StatusCreated, wt)
}
