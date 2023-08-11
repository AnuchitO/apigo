package wallet

import (
	"database/sql"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB
var once sync.Once

func init() {
	Conn()
	createTable()
}

func Conn() *sql.DB {
	// TODO: DATABASE_URL from env
	var url string = "postgres://batsjuib:Voopy92tnjyUMBQhi0EDtxTdg--aA-rK@tiny.db.elephantsql.com/batsjuib"
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", url)
		if err != nil {
			log.Fatal("can not connect to database:", err)
		}

		if err := db.Ping(); err != nil {
			log.Fatal(err)
		}
	})

	log.Println("connected to database")

	return db
}

const createWalletTable = `
CREATE TABLE IF NOT EXISTS wallets (
	id SERIAL PRIMARY KEY,
	owner TEXT,
	balance FLOAT
);
`

func createTable() {
	conn := Conn()
	_, err := conn.Exec(createWalletTable)
	if err != nil {
		log.Fatal(err)
	}
}

const insertWallet = `
	INSERT INTO wallets (owner, balance)
	VALUES ($1, $2)
	RETURNING id, owner, balance;
`

const getWalletByID = `
	SELECT id, owner, balance
	FROM wallets
	WHERE id = $1;
`

const getBalanceByID = `
	SELECT owner, balance
	FROM wallets
	WHERE id = $1;
`
const depositByID = `
	UPDATE wallets
	SET balance = balance + $1
	WHERE id = $2
	RETURNING id, owner, balance;
`

const withdrawByID = `
	UPDATE wallets
	SET balance = balance - $1
	WHERE id = $2
	RETURNING id, owner, balance;
`

func WithdrawByIDHandler(c *gin.Context) {
	id := convertID(c.Param("id"))

	conn := Conn()

	var wt Wallet
	err := conn.QueryRow(getWalletByID, id).Scan(&wt.ID, &wt.Owner, &wt.Balance)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "wallet not found",
		})
		return
	}

	var amount struct {
		Amount float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&amount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if wt.Balance < amount.Amount {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "balance not enough",
		})
		return
	}

	wt.Balance -= amount.Amount

	err = db.QueryRow(withdrawByID, amount.Amount, id).Scan(&wt.ID, &wt.Owner, &wt.Balance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      wt.ID,
		"balance": wt.Balance,
	})
}

func DepositByIDHandler(c *gin.Context) {
	id := convertID(c.Param("id"))

	conn := Conn()

	var wt Wallet
	err := conn.QueryRow(getWalletByID, id).Scan(&wt.ID, &wt.Owner, &wt.Balance)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "wallet not found",
		})
		return
	}

	var amount struct {
		Amount float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&amount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	am := wt.Balance + amount.Amount

	err = conn.QueryRow(depositByID, am, id).Scan(&wt.ID, &wt.Owner, &wt.Balance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      wt.ID,
		"balance": wt.Balance,
	})
}

func GetWalletByIDHandler(c *gin.Context) {
	id := convertID(c.Param("id"))

	conn := Conn()

	var wt Wallet
	err := conn.QueryRow(getWalletByID, id).Scan(&wt.ID, &wt.Owner, &wt.Balance)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "wallet not found",
		})
		return
	}

	c.JSON(http.StatusOK, wt)
}

func GetBalanceByIDHandler(c *gin.Context) {
	id := convertID(c.Param("id"))

	conn := Conn()

	var balance float64
	var owner string
	err := conn.QueryRow(getBalanceByID, id).Scan(&owner, &balance)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "wallet not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"owner":   owner,
		"balance": balance,
	})
}
