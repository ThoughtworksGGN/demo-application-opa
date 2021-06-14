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

func(s *Store) FindTicketsByAsignee(assignee uuid.UUID) []Ticket {
	var tickets []Ticket

	for _, ticket := range s.Tickets {
		if ticket.Assignee == assignee {
			tickets = append(tickets, ticket)
		}
	}

	return tickets
}
