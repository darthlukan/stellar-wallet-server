package api

import (
	"net/http"

	"github.com/darthlukan/stellar-wallet-server/account"
	"github.com/darthlukan/stellar-wallet-server/payments"
	"github.com/gin-gonic/gin"
)

const (
	environ string = "test"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "PONG!"})
	return
}

func Health(c *gin.Context) {
	data := gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"mode": gin.Mode(),
		},
	}
	c.JSON(http.StatusOK, data)
	return
}

func CreateAccount(c *gin.Context) {
	account.CreateAccount(c)
	return
}

func GetAccount(c *gin.Context) {
	address := c.Param("address")
	account.GetAccount(address, c)
	return
}

func SendPayment(c *gin.Context) {
	var data map[string]interface{}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"data":   nil,
			"error":  err,
		})
		return
	}

	srcSecKey := data["seed"].(string)
	destAddr := data["dest"].(string)
	amount := data["amount"].(string)
	assetType := data["asset"].(string)

	txn, err := payments.SendPayment(environ, srcSecKey, destAddr, amount, assetType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"data":   nil,
			"error":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   txn,
		"error":  nil,
	})
	return

}
