package payments

import (
	"log"
	"testing"

	"github.com/darthlukan/stellar-wallet-server/account"
	"github.com/darthlukan/stellar-wallet-server/harness"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/txnbuild"
	"github.com/stretchr/testify/assert"
)

func TestBuildPaymentTransaction(t *testing.T) {
	testSrcAccount, err := harness.CreateTestAccount()
	if err != nil {
		t.Fatalf("harness.CreateTestAccount():err = %v; want nil", err)
	}

	_, er := harness.FundTestAccount(&testSrcAccount)
	if er != nil {
		t.Fatalf("harness.FundTestAccount(&testSrcAccount):err = %v; want nil", er)
	}

	testDestAccount, err := harness.CreateTestAccount()
	if err != nil {
		t.Fatalf("harness.CreateTestAccount():err = %v; want nil", err)
	}

	_, e := harness.FundTestAccount(&testDestAccount)
	if e != nil {
		t.Fatalf("harness.FundTestAccount(&testSrcAccount):err = %v; want nil", e)
	}

	destAddr := testDestAccount.Address
	testAmount := "5"
	baseFee := txnbuild.MinBaseFee
	timeout := txnbuild.NewInfiniteTimeout()
	asset := txnbuild.NativeAsset{}

	srcAccount, err := account.QueryAccount("test", testSrcAccount.Address)
	if err != nil {
		t.Fatalf("account.QueryAccount(testSrcAccount.Address):err = %v; want nil", err)
	}

	txn, err := BuildPaymentTransaction(
		srcAccount,
		destAddr,
		testAmount,
		baseFee,
		timeout,
		asset)

	if err != nil {
		t.Fatalf("BuildPaymentTransaction:err = %v; want nil", err)
	}

	assert.Equalf(t, txn.SourceAccount().AccountID, srcAccount.AccountID, "should be equal")
	assert.Greaterf(t, txn.MaxFee(), int64(0), "should be greater")
	assert.Greaterf(t, len(txn.Operations()), 0, "should be greater")

	srcKeyPair := keypair.MustParseFull(testSrcAccount.Seed)

	t.Run("TestSignTransaction", func(t *testing.T) {
		tx, err := SignTransaction("test", srcKeyPair, txn)
		if err != nil {
			t.Fatalf("SignTransaction('test', srcKeyPair, txn):err = %v; want nil", err)
		}
		assert.Greaterf(t, len(tx.Signatures()), 0, "should be greater")
	})

	t.Run("TestSendTransaction", func(t *testing.T) {
		tx, err := SendPayment("test", srcKeyPair.Seed(), destAddr, testAmount, "XLM")
		if err != nil {
			t.Fatalf("SendPayment:err = %v; want nil", err)
		}
		log.Printf("tx is %T; tx = %v;", tx, tx)
		assert.NotEmptyf(t, tx.AccountSequence, "Should not be empty")
	})
}

// This is covered by a nested test run, no need to do it twice
func TestSignTransaction(t *testing.T) {}

// Also covered by nested test run
func TestSendPayment(t *testing.T) {}
