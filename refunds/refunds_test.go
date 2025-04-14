package refunds

import (
	"context"
	"github.com/mdwt/paystack-go/client"
	"github.com/mdwt/paystack-go/common"
	"github.com/mdwt/paystack-go/logger"
	"testing"
)

func TestInitializeTransaction(t *testing.T) {
	txn := &CreateRefundRequest{
		Transaction: "456484564",
	}

	refunds := New(client.Options{
		ApiKey:    "sk_test_e39ce23869e6e677121a5e6ef691a8c3d835f0bb",
		ConnectId: "",
		BaseUrl:   common.BaseURLV1,
	}, logger.NewDefaultLogger())

	_, err := refunds.Create(context.Background(), txn)
	if err != nil {
		t.Error(err)
	}

}
