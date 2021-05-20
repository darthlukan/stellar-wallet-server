package lib

import (
	"testing"

	"github.com/stellar/go/network"
	"github.com/stretchr/testify/assert"
)

const (
	TestEnv string = "test"
	ProdEnv string = "prod"
	FailEnv string = "fail"
)

// Stub, we don't need to test the stellar go SDK
func TestCreateKeyPair(t *testing.T) {}

func TestGetHorizonClient(t *testing.T) {
	testClient := GetHorizonClient(TestEnv)
	assert.Containsf(t, testClient.HorizonURL, TestEnv, "should contain 'test'")

	prodClient := GetHorizonClient(ProdEnv)
	assert.NotContainsf(t, prodClient.HorizonURL, TestEnv, "should not contain 'test'")

	failClient := GetHorizonClient(FailEnv)
	assert.Containsf(t, failClient.HorizonURL, TestEnv, "shoujld contain 'test'")
}

func TestGetNetworkPassphrase(t *testing.T) {
	testPass := GetNetworkPassphrase(TestEnv)
	// We should get the network.TestNetworkPassphrase
	assert.Equalf(t, testPass, network.TestNetworkPassphrase, "should be equal")

	prodPass := GetNetworkPassphrase(ProdEnv)
	// We should get network.PublicNetworkPassphrase
	assert.Equalf(t, prodPass, network.PublicNetworkPassphrase, "should be equal")

	failPass := GetNetworkPassphrase(FailEnv)
	// We should get our default, network.TestNetworkPassphrase
	assert.Equalf(t, failPass, network.TestNetworkPassphrase, "should be equal")
}
