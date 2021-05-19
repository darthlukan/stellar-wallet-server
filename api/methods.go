package api

import (
	"net/http"

	"github.com/darthlukan/stellar-wallet-server/account"
	"github.com/gin-gonic/gin"
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
