package wallet

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertWallet(wt Wallet) (Wallet, error) {
	conn := Conn()
	var w Wallet
	err := conn.QueryRow(insertWallet, wt.Owner, wt.Balance).Scan(&w.ID, &w.Owner, &w.Balance)
	if err != nil {
		return Wallet{}, err
	}

	return w, nil
}

func CreateWalletHandler(c *gin.Context) {
	var wt Wallet
	if err := c.ShouldBindJSON(&wt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	wt, err := InsertWallet(wt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, wt)
}
