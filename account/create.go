package account

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stellar/go/keypair"
)

func createKeyPair() (string, string, error) {
	pair, err := keypair.Random()
	return pair.Address(), pair.Seed(), err
}

func CreateKeyPair() (string, string, error) {
	return createKeyPair()
}

func CreateAccount(c *gin.Context) {
	address, seed, err := createKeyPair()
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
