package domain

import apperrors "github.com/jmechavez/user-management-api/internal/appErrors"

type User struct {
	IdNumber  int64  `json:"idNumber,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
}

type UserRepository interface {
	FindAll() ([]User, *apperrors.AppError)
	FindAllv2() ([]User, *apperrors.AppError)
	FindByID(id int64) (*User, *apperrors.AppError)
}
