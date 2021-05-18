package account

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	hClient "github.com/stellar/go/clients/horizonclient"
	hprotocol "github.com/stellar/go/protocols/horizon"
)

func getAccount(address string) hprotocol.Account {
	client := hClient.DefaultTestNetClient
	request := hClient.AccountRequest{AccountID: address}
	account, err := client.AccountDetail(request)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return account
}

func GetAccount(address string, c *gin.Context) {
	account := getAccount(address)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"account": account,
	})
	return
}
