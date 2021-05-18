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
		panic(err)
	}

	client := hClient.DefaultTestNetClient
	accountRequest := hClient.AccountRequest{AccountID: address}
	account, err := client.AccountDetail(accountRequest)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"account": account,
		"seed":    seed,
	})

	return
}
