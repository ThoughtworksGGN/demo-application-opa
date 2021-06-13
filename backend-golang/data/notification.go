package data

import uuid "github.com/satori/go.uuid"

type Notification struct {
	NotificationId uuid.UUID `json:"notificationid"`
	AccountId      uuid.UUID `json:"accountid"`
	Info           string    `json:"info"`
}

func dummyNotifications() []Notification {
	return []Notification{
		{
			NotificationId: uuid.FromStringOrNil("08718203-E617-4DA5-85DE-831501D6CE84"),
			AccountId:      uuid.FromStringOrNil("9B8FF595-9E32-41CA-8CF1-068CAC73A0B0"),
			Info:           "Notification1 for User1",
		},
		{
			NotificationId: uuid.FromStringOrNil("CBD1B44D-27A3-4713-9B86-558CB1100AF0"),
			AccountId:      uuid.FromStringOrNil("9B8FF595-9E32-41CA-8CF1-068CAC73A0B0"),
			Info:           "Notification2 for User1",
		},
		{
			NotificationId: uuid.FromStringOrNil("C811051B-3108-4C5D-9BB9-F3830C008CB5"),
			AccountId:      uuid.FromStringOrNil("FAB1D671-2402-4CB1-B295-ECF00D4C4065"),
			Info:           "Notification1 for User2",
		},
		{
			NotificationId: uuid.FromStringOrNil("C4DFE41E-9F17-4DBB-96C3-4CB50B3E948E"),
			AccountId:      uuid.FromStringOrNil("FAB1D671-2402-4CB1-B295-ECF00D4C4065"),
			Info:           "Notification2 for User2",
		},
	}
}
