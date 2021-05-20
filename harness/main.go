package harness

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/darthlukan/stellar-wallet-server/lib"
)

const (
	friendBotUrl string = "https://friendbot.stellar.org/?addr="
)

type TestAccount struct {
	Address string
	Seed    string
}

func CreateTestAccount() (TestAccount, error) {
	address, seed, err := lib.CreateKeyPair()
	if err != nil {
		log.Printf("account.CreateKeyPair():err = %v; want nil", err)
	}

	taccount := TestAccount{
		Address: address,
		Seed:    seed,
	}
	return taccount, err
}

func FundTestAccount(account *TestAccount) (string, error) {
	address := account.Address
	resp, err := http.Get(friendBotUrl + address)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	accountJson := string(body)
	return accountJson, nil
}
