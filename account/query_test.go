package account

import (
	"testing"

	"github.com/darthlukan/stellar-wallet-server/harness"
	"github.com/stretchr/testify/assert"
)

func TestGetAccount(t *testing.T) {
	// "TEST" is not a valid address, so we expect failure
	account, err := QueryAccount("test", "TEST")
	if len(account.AccountID) > 0 {
		t.Errorf("getAccount('TEST'):account.AccountID = %v; want ''", account.AccountID)
	}
	if err == nil {
		t.Errorf("getAccount('TEST'):err = %v; want 'horizon error: \"Bad Request\" - check horizon.Error.Problem for more information'", err)
	}

	taccount, err := harness.CreateTestAccount()
	if err != nil {
		t.Fatalf("harness.CreateTestAccount():err = %v; want nil", err)
	}

	// We don't care about the returned JSON in this case,
	// if err is nil we're funded and can operate
	_, er := harness.FundTestAccount(&taccount)
	if er != nil {
		t.Fatalf("harness.FundTestAccount(&taccount):err = %v; want nil", er)
	}

	rtAccount, err := QueryAccount("test", taccount.Address)
	if err != nil {
		t.Fatalf("getAccount(taccount.Address):err = %v; want nil", err)
	}

	assert.Greater(t, len(rtAccount.Balances), 0, "We should have at least one balance")
	assert.NotEmpty(t, rtAccount.AccountID, "AccountID should not be empty string")
}
