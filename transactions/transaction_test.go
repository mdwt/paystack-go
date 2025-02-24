package transactions

import (
	"context"
	"fmt"
	"github.com/mdwt/paystack-go"
	"math/rand"
	"testing"
)

func TestInitializeTransaction(t *testing.T) {
	txn := &TransactionRequest{
		Email:     "user123@gmail.com",
		Amount:    6000,
		Reference: fmt.Sprintf("ref_%d", rand.Intn(1000000)),
	}

	transaction := Client{
		client: paystack.getClient("sk_test_e39ce23869e6e677121a5e6ef691a8c3d835f0bb"),
	}

	_, err := transaction.Initialize(context.Background(), txn)
	if err != nil {
		t.Error(err)
	}

}

func TestChargeAuthorization(t *testing.T) {
	txn := ChargeAuthorizationRequest{
		Email:             "koos@test.com",
		Amount:            6000,
		AuthorizationCode: "AUTH_468hvddfw6",
		Currency:          "ZAR",
		Reference:         fmt.Sprintf("ref_%d", rand.Intn(1000000)),
	}

	transaction := Client{
		client: paystack.getClient("sk_test_e39ce23869e6e677121a5e6ef691a8c3d835f0bb"),
	}

	_, err := transaction.ChargeAuthorization(context.Background(), txn)
	if err != nil {
		t.Error(err)
	}
}
