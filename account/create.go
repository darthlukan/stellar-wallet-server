package account

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	hClient "github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
)

func createKeyPair() (string, string, error) {
	pair, err := keypair.Random()
	return pair.Address(), pair.Seed(), err
}

func CreateAccount(c *gin.Context) {
	address, seed, err := createKeyPair()
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"account": nil,
			"seed":    nil,
			"error":   err,
		})
		return
	}

	client := hClient.DefaultTestNetClient
	accountRequest := hClient.AccountRequest{AccountID: address}
	account, err := client.AccountDetail(accountRequest)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"account": nil,
			"seed":    nil,
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"account": account,
		"seed":    seed,
		"error":   nil,
	})

	return
}
