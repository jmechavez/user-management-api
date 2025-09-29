package domain

type UserRepositoryStub struct {
	users []User
}

func (r UserRepositoryStub) FindAll() ([]User, error) {
	return r.users, nil
}

func NewUserRepositoryStub() UserRepositoryStub {
	users := []User{
		{
			IdNumber:  123,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johndoe@test.com",
		},
	}
	return UserRepositoryStub{users: users}
}
