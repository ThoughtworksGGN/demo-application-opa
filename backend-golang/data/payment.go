package data

import uuid "github.com/satori/go.uuid"

type Payment struct {
	PaymentId uuid.UUID `json:"paymentid"`
	AccountId uuid.UUID `json:"accountid"`
	Amount    int64     `json:"amount"`
}

func dummyPayments() []Payment {
	return []Payment{
		{
			PaymentId: uuid.FromStringOrNil("85BBDE19-1409-4A10-94F3-A7F79874304E"),
			AccountId: uuid.FromStringOrNil("9B8FF595-9E32-41CA-8CF1-068CAC73A0B0"),
			Amount:    int64(100),
		},
		{
			PaymentId: uuid.FromStringOrNil("563E8CD0-9F35-47D3-9F38-4980FCDF9583"),
			AccountId: uuid.FromStringOrNil("FAB1D671-2402-4CB1-B295-ECF00D4C4065"),
			Amount:    int64(100),
		},
		{
			PaymentId: uuid.FromStringOrNil("2AD8904F-566E-4872-BDB8-8783F14910D3"),
			AccountId: uuid.FromStringOrNil("9B8FF595-9E32-41CA-8CF1-068CAC73A0B0"),
			Amount:    int64(100),
		},
		{
			PaymentId: uuid.FromStringOrNil("1D42A614-8F06-4F06-88EC-589D4A8D3118"),
			AccountId: uuid.FromStringOrNil("9B8FF595-9E32-41CA-8CF1-068CAC73A0B0"),
			Amount:    int64(100),
		},
	}
}
