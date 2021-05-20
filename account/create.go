package account

import (
	"log"
	"net/http"

	"github.com/darthlukan/stellar-wallet-server/lib"
	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	address, seed, err := lib.CreateKeyPair()
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"account": "",
			"seed":    "",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"address": address,
		"seed":    seed,
		"error":   err,
	})
	return
}
