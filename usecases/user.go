package usecases

import "project_university/domain"

type User interface {
	Get(user *domain.User) (*domain.User, error)
	Post(user *domain.User) error
	Put(user *domain.User) error
	Delete(user *domain.User) error
}
