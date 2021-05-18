package account

import (
	"testing"
)

func TestGetAccount(t *testing.T) {
	// "TEST" is not a valid address, so we expect failure
	account, err := getAccount("TEST")
	if len(account.AccountID) > 0 {
		t.Errorf("getAccount('TEST'):account.AccountID = %v; want ''", account.AccountID)
	}
	if err == nil {
		t.Errorf("getAccount('TEST'):err = %v; want 'horizon error: \"Bad Request\" - check horizon.Error.Problem for more information'", err)
	}
}
