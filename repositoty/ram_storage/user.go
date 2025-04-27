package ramstorage

import "project_university/domain"

type User struct {
	data map[string]*domain.User
}

func NewUser() *User {
	return &User{
		data: make(map[string]*domain.User),
	}
}

func (u *User) Get(user *domain.User) (*domain.User, error) {
	if _, err := u.data[user.Login]; !err {
		return nil, domain.UserNotFound
	}
	return u.data[user.Login], nil
}

func (u *User) Post(user *domain.User) error {
	if _, err := u.data[user.Login]; err {
		return domain.UserAlreadyExists
	}
	u.data[user.Login] = user
	return nil
}

func (u *User) Put(user *domain.User) error {
	if _, err := u.data[user.Login]; !err {
		return domain.UserNotFound
	}
	u.data[user.Login] = user
	return nil
}

func (u *User) Delete(user *domain.User) error {
	if _, err := u.data[user.Login]; !err {
		return domain.UserNotFound
	}
	delete(u.data, user.Login)
	return nil
}
