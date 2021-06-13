package data

import uuid "github.com/satori/go.uuid"

type User struct {
	UserId   uuid.UUID `json:"userid"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
}

const CUSTOMER = "CUSTOMER"
const SUPPORT = "SUPPORT"

func dummyUsers() []User {
	return []User{
		{
			UserId: uuid.FromStringOrNil("2F6AEE39-AA31-421F-9396-4D7066BBBB74"),
			Name: "user1",
			Password: "password1",
			Role: CUSTOMER,
		},
		{
			UserId: uuid.FromStringOrNil("C140BE51-C381-47BC-8452-7C2F7E49E4DE"),
			Name: "user2",
			Password: "password2",
			Role: CUSTOMER,
		},
		{
			UserId: uuid.FromStringOrNil("423C3C1E-3188-4471-BADA-C0EF06DD6419"),
			Name: "support1",
			Password: "password3",
			Role: SUPPORT,
		},
		{
			UserId: uuid.FromStringOrNil("6A22DAF5-E012-402A-B61F-9C84CB5BE0D1"),
			Name: "support2",
			Password: "password4",
			Role: SUPPORT,
		},
	}
}