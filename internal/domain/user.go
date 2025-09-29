package domain

type User struct {
	IdNumber  int64
	FirstName string
	LastName  string
	Email     string
}

type UserRepository interface {
	FindAll() ([]User, error)
}
