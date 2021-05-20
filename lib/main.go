package lib

import (
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
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

func GetNetworkPassphrase(environ string) string {
	switch environ {
	case "test":
		networkPassphrase := network.TestNetworkPassphrase
	case "prod":
		networkPassphrase := network.PublicNetworkPassphrase
	default:
		networkPassphrase := network.TestNetworkPassphrase
	}
	return networkPassphrase
}

func CreateKeyPair() (string, string, error) {
	pair, err := keypair.Random()
	return pair.Address(), pair.Seed(), err
}
