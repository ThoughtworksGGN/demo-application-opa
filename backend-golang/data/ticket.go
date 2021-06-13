package data

import uuid "github.com/satori/go.uuid"

type Ticket struct {
	TicketId  uuid.UUID `json:"ticketid"`
	AccountId uuid.UUID `json:"accountid"`
	Assignee  uuid.UUID `json:"assignee"`
	Info      string    `json:"info"`
	Notes     string    `json:"notes"`
}

func dummyTickets() []Ticket {
	return []Ticket{
		{
			TicketId:  uuid.FromStringOrNil("96E20F29-D700-4638-9636-77E302F78B51"),
			AccountId: uuid.FromStringOrNil("9B8FF595-9E32-41CA-8CF1-068CAC73A0B0"),
			Assignee:  uuid.FromStringOrNil("423C3C1E-3188-4471-BADA-C0EF06DD6419"),
			Info:      "It doesnt work",
		},
		{
			TicketId:  uuid.FromStringOrNil("4E79C245-87AC-47D0-87A4-C4BABCBA9847"),
			AccountId: uuid.FromStringOrNil("9B8FF595-9E32-41CA-8CF1-068CAC73A0B0"),
			Assignee:  uuid.FromStringOrNil("423C3C1E-3188-4471-BADA-C0EF06DD6419"),
			Info:      "Please fix it",
		},
		{
			TicketId:  uuid.FromStringOrNil("2C5268F9-17CB-465B-B8ED-48307165E0AF"),
			AccountId: uuid.FromStringOrNil("FAB1D671-2402-4CB1-B295-ECF00D4C4065"),
			Assignee:  uuid.FromStringOrNil("423C3C1E-3188-4471-BADA-C0EF06DD6419"),
			Info:      "I pay money. Gib MOvie",
		},
		{
			TicketId:  uuid.FromStringOrNil("DCC3F319-A684-4409-AD5C-1FA3A2CCF107"),
			AccountId: uuid.FromStringOrNil("FAB1D671-2402-4CB1-B295-ECF00D4C4065"),
			Assignee:  uuid.FromStringOrNil("6A22DAF5-E012-402A-B61F-9C84CB5BE0D1"),
			Info:      "Worst service. Used movie given. Replace with new movie",
		},
	}
}
