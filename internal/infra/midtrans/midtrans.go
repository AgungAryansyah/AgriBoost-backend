package midtrans

import (
	"AgriBoost/internal/infra/env"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransItf interface {
	NewTransactionToken(orderId string) (*snap.Response, error)
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

func (m *Midtrans) NewTransactionToken(orderId string) (*snap.Response, error) {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: 100000,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}

	snapResp, err := m.Client.CreateTransaction(req)
	return snapResp, err
}
