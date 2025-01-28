package paystack

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestInitializeTransaction(t *testing.T) {
	txn := &TransactionRequest{
		Email:     "user123@gmail.com",
		Amount:    6000,
		Reference: fmt.Sprintf("ref_%d", rand.Intn(1000000)),
	}

	transaction := TransactionService{
		client: getClient("sk_test_e39ce23869e6e677121a5e6ef691a8c3d835f0bb"),
	}

	_, err := transaction.Initialize(txn)
	if err != nil {
		t.Error(err)
	}

}
