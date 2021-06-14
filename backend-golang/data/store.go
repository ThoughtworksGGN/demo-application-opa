package data

type Store struct {
	Users         []User         `json:"users"`
	Accounts      []Account      `json:"accounts"`
	Payments      []Payment      `json:"payments"`
	Notifications []Notification `json:"notifications"`
	Tickets       []Ticket       `json:"tickets"`
	Sessions      []string       `json:"sessions"`
}

type UserNotFoundError struct{}

func (s *Store) Initialise() {
	s.Users = dummyUsers()
	s.Accounts = dummyAccounts()
	s.Payments = dummyPayments()
	s.Notifications = dummyNotifications()
	s.Tickets = dummyTickets()
	s.Sessions = []string{}
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

func (e UserNotFoundError) Error() string {
	return "User Not found"
}
