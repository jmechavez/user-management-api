package domain

type User struct {
	IdNumber  int64  `json:"idNumber,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
}

type UserRepository interface {
	FindAll() ([]User, error)
	FindAllv2() ([]User, error)
}
