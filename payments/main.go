package payments

import (
	"log"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/protocols/horizon"
	"github.com/stellar/go/txnbuild"

	"github.com/darthlukan/stellar-wallet-server/lib"
)

func BuildPaymentTransaction(
	srcAccount horizon.Account,
	destAddr string,
	amount string,
	baseFee int,
	timeout txnbuild.Timebounds,
	asset txnbuild.Asset) (*txnbuild.Transaction, error) {

	txn, err := txnbuild.NewTransaction(
		txnbuild.TransactionParams{
			SourceAccount:        &srcAccount,
			IncrementSequenceNum: true,
			BaseFee:              int64(baseFee),
			Timebounds:           timeout,
			Operations: []txnbuild.Operation{
				&txnbuild.Payment{
					Destination: destAddr,
					Amount:      amount,
					Asset:       asset,
				},
			},
		},
	)
	return txn, err
}

func SignTransaction(environ string, srcKeyPair *keypair.Full, txn *txnbuild.Transaction) (*txnbuild.Transaction, error) {
	netPass := lib.GetNetworkPassphrase(environ)
	txn, err := txn.Sign(netPass, srcKeyPair)
	return txn, err
}

func SendPayment(environ, srcSecKey, destAddr, amount, assetType string) (horizon.Transaction, error) {
	client := lib.GetHorizonClient(environ)

	destAccRequest := horizonclient.AccountRequest{AccountID: destAddr}
	destAccount, err := client.AccountDetail(destAccRequest)
	if err != nil {
		// TODO: Bubble up a "destination account does not exist or is invalid" error
		log.Panic(err)
	}

	srcKeyPair := keypair.MustParseFull(srcSecKey)
	srcAccRequest := horizonclient.AccountRequest{AccountID: srcKeyPair.Address()}
	srcAccount, err := client.AccountDetail(srcAccRequest)
	if err != nil {
		// TODO: Bubble up a "Your account is not funded or does not exist" error
		log.Panic(err)
	}

	tx, err := BuildPaymentTransaction(
		srcAccount,
		destAccount.AccountID,
		amount,
		txnbuild.MinBaseFee,
		txnbuild.NewInfiniteTimeout(),
		txnbuild.NativeAsset{})

	if err != nil {
		log.Panic(err)
	}

	txn, err := SignTransaction(environ, srcKeyPair, tx)
	if err != nil {
		log.Panic(err)
	}

	resp, err := client.SubmitTransaction(txn)
	if err != nil {
		log.Panic(err)
	}
	return resp, err
}
