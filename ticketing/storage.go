package ticketing

// Storage ...
type Storage interface {
	GetUsers(p Pagination, criteria UsersSearchCriteria) ([]User, error)
	FindUser(id int) (User, error)
	CreateUser(user User) (User, error)
}

type Pagination struct {
	Limit int
	Page  int
}
type UsersSearchCriteria struct {
	Name string
}
