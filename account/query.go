package account

import (
	"log"
	"net/http"

	"github.com/darthlukan/stellar-wallet-server/lib"
	"github.com/gin-gonic/gin"
	hClient "github.com/stellar/go/clients/horizonclient"
	hprotocol "github.com/stellar/go/protocols/horizon"
)

func QueryAccount(environ, address string) (hprotocol.Account, error) {
	client := lib.GetHorizonClient(environ)
	request := hClient.AccountRequest{AccountID: address}
	account, err := client.AccountDetail(request)

	return account, err
}

func GetAccount(address string, c *gin.Context) {
	account, err := QueryAccount("prod", address)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"account": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"account": account,
	})
	return
}
