package services

import (
	apperrors "github.com/jmechavez/user-management-api/internal/appErrors"
	"github.com/jmechavez/user-management-api/internal/domain"
)

type UserServices interface {
	FindAllUser() ([]domain.User, *apperrors.AppError)
	FindById(int64) (*domain.User, *apperrors.AppError)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func (s DefaultUserService) FindAllUser() ([]domain.User, *apperrors.AppError) {
	// return s.repo.FindAll()
	return s.repo.FindAll()
}

func (s DefaultUserService) FindById(id int64) (*domain.User, *apperrors.AppError) {
	return s.repo.FindByID(id)
}

func NewUserService(repository domain.UserRepository) DefaultUserService {
	return DefaultUserService{repository}
}
