package data

import uuid "github.com/satori/go.uuid"

type Store struct {
	Users         []User          `json:"users"`
	Accounts      []Account       `json:"accounts"`
	Payments      []Payment       `json:"payments"`
	Notifications []Notification  `json:"notifications"`
	Tickets       []Ticket        `json:"tickets"`
	Sessions      map[string]User `json:"sessions"`
}

type RegoData struct {
	Users map[uuid.UUID]User `json:"users"`
	Accounts map[uuid.UUID]Account `json:"accounts"`
	Tickets map[uuid.UUID]Ticket `json:"tickets"`
}

type UserNotFoundError struct{}

func (s *Store) Initialise() {
	s.Users = dummyUsers()
	s.Accounts = dummyAccounts()
	s.Payments = dummyPayments()
	s.Notifications = dummyNotifications()
	s.Tickets = dummyTickets()
	s.Sessions = make(map[string]User)
}

func (s *Store) FindUserByName(name string) (User, error) {
	users := s.Users
	userFound := false
	var matchingUser User

	for _, user := range users {
		if user.Name == name {
			matchingUser = user
			userFound = true
			break
		}
	}

	if userFound {
		return matchingUser, nil
	}
	return matchingUser, &UserNotFoundError{}
}

func (s *Store) SessionExists(token string) bool {
	sessionExists := false
	sessions := s.Sessions

	if _, ok := sessions[token]; ok {
		sessionExists = true
	}

	return sessionExists
}

func (e UserNotFoundError) Error() string {
	return "User Not found"
}

func (s *Store) FindTicketsByAsignee(assignee uuid.UUID) []Ticket {
	var tickets []Ticket

	for _, ticket := range s.Tickets {
		if ticket.Assignee == assignee {
			tickets = append(tickets, ticket)
		}
	}

	return tickets
}

func (s *Store) FindTicketsByAccountId(accountId uuid.UUID)[]Ticket{
	var tickets []Ticket

	for _, ticket := range s.Tickets {
		if ticket.AccountId == accountId {
			tickets = append(tickets, ticket)
		}
	}

	return tickets
}

func (s *Store) FindPaymentsByAccountId(accountId uuid.UUID)[]Payment{
	var payments []Payment

	for _, payment := range s.Payments {
		if payment.AccountId == accountId {
			payments = append(payments, payment)
		}
	}

	return payments
}

func (s *Store) FindNotificationsByAccountId(accountId uuid.UUID)[]Notification{
	var notifications []Notification

	for _, notification := range s.Notifications {
		if notification.AccountId == accountId {
			notifications = append(notifications, notification)
		}
	}

	return notifications
}

func (s *Store) FindUserByUserId(id uuid.UUID) (User, error) {
	users := s.Users
	userFound := false
	var matchingUser User

	for _, user := range users {
		if user.UserId == id {
			matchingUser = user
			userFound = true
			break
		}
	}

	if userFound {
		return matchingUser, nil
	}
	return matchingUser, &UserNotFoundError{}
}

func (s *Store) RegoData() RegoData {
	data := RegoData{
		Users: make(map[uuid.UUID]User),
		Accounts: make(map[uuid.UUID]Account),
		Tickets: make(map[uuid.UUID]Ticket),
	}

	for _, user := range s.Users {
		data.Users[user.UserId] = user
	}

	for _, account := range s.Accounts {
		data.Accounts[account.AccountId] = account
	}

	for _, ticket := range s.Tickets {
		data.Tickets[ticket.TicketId] = ticket
	}
	return data
}