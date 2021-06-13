package data

type Store struct {
	Users         []User
	Accounts      []Account
	Payments      []Payment
	Notifications []Notification
	Tickets       []Ticket
	Sessions      []string
}

func (s *Store) Initialise() {
	s.Users = dummyUsers()
	s.Accounts = dummyAccounts()
	s.Payments = dummyPayments()
	s.Notifications = dummyNotifications()
	s.Tickets = dummyTickets()
	s.Sessions = []string{}
}
