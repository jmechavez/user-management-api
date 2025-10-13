package services

import "github.com/jmechavez/user-management-api/internal/domain"

type UserServices interface {
	FindAllUser() ([]domain.User, error)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func (s DefaultUserService) FindAllUser() ([]domain.User, error) {
	return s.repo.FindAll()
}

func NewUserService(repository domain.UserRepository) DefaultUserService {
	return DefaultUserService{repository}
}
