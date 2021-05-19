package harness

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/darthlukan/stellar-wallet-server/account"
)

const (
	friendBotUrl string = "https://friendbot.stellar.org/?addr="
)

type TestAccount struct {
	Address string
	Seed    string
}

func CreateTestAccount() (TestAccount, error) {
	address, seed, err := account.CreateKeyPair()
	if err != nil {
		log.Printf("account.CreateKeyPair():err = %v; want nil", err)
	}

	account := TestAccount{
		Address: address,
		Seed:    seed,
	}
	return TestAccount, err
}

func FundTestAccount(account *TestAccount) error {
	address := account.Address
	resp, err := http.Get(friendBotUrl + address)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("body is %T; body = %v", body, string(body))
	return nil
}
