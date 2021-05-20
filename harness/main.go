package harness

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/stellar/go/keypair"
)

const (
	friendBotUrl string = "https://friendbot.stellar.org/?addr="
)

type TestAccount struct {
	Address string
	Seed    string
}

func createKeyPair() (string, string, error) {
	pair, err := keypair.Random()
	return pair.Address(), pair.Seed(), err
}

func CreateTestAccount() (TestAccount, error) {
	address, seed, err := createKeyPair()
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
