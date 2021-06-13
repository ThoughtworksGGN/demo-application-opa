package data

import uuid "github.com/satori/go.uuid"

type Account struct {
	AccountId uuid.UUID `json:"accountid"`
	UserId    uuid.UUID `json:"userid"`
}

func dummyAccounts() []Account{
	return []Account{
		{
			AccountId: uuid.FromStringOrNil("9B8FF595-9E32-41CA-8CF1-068CAC73A0B0"),
			UserId: uuid.FromStringOrNil("2F6AEE39-AA31-421F-9396-4D7066BBBB74"),
		},
		{
			AccountId: uuid.FromStringOrNil("FAB1D671-2402-4CB1-B295-ECF00D4C4065"),
			UserId: uuid.FromStringOrNil("C140BE51-C381-47BC-8452-7C2F7E49E4DE"),
		},
	}
}