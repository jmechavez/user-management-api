package domain

type User struct {
	IdNumber  int64  `json:"idNumber"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type UserRepository interface {
	FindAll() ([]User, error)
}
