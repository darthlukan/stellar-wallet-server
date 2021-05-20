package lib

import "github.com/stellar/go/keypair"

func CreateKeyPair() (string, string, error) {
	pair, err := keypair.Random()
	return pair.Address(), pair.Seed(), err
}
