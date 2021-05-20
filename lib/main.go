package lib

import (
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
)

func GetHorizonClient(clientType string) horizonclient.Client {
	switch clientType {
	case "test":
		client := horizonclient.DefaultTestNetClient
	case "prod":
		client := horizonclient.DefaultPublicNetClient
	default:
		client := horizonclient.DefaultTestNetClient
	}
	return client
}

func CreateKeyPair() (string, string, error) {
	pair, err := keypair.Random()
	return pair.Address(), pair.Seed(), err
}
