package midtrans

import (
	"AgriBoost/internal/infra/env"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransItf interface {
	NewTransactionToken(orderId string, amount int64) (string, error)
}

type Midtrans struct {
	Client snap.Client
}

func NewMidtrans(env env.Env) MidtransItf {
	client := snap.Client{}
	client.New(env.MIDTRANS_SERVER_KEY, midtrans.Sandbox)

	return &Midtrans{
		Client: client,
	}
}

func (m *Midtrans) NewTransactionToken(orderId string, amount int64) (string, error) {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: amount,
		},
	}

	snapUrl, err := m.Client.CreateTransactionUrl(req)
	return snapUrl, err
}
