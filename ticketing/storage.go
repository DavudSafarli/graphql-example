package ticketing

// Storage ...
type Storage interface {
	GetUsers(p Pagination, criteria UsersSearchCriteria) ([]User, error)
	FindUser(id int) (User, error)
	CreateUser(user User) (User, error)

	GetTicketsAssignees(id []int) (map[int][]User, error)
	GetTicketAssignees(id int) ([]User, error)

	GetTickets(p Pagination) ([]Ticket, error)
	CreateTicket(ticket Ticket) (Ticket, error)
	FindTicket(id int) (Ticket, error)
}

type Pagination struct {
	Limit int
	Page  int
}
type UsersSearchCriteria struct {
	Name string
}
