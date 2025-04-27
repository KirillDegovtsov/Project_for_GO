package service

import (
	"project_university/domain"
	"project_university/repositoty"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	repo repositoty.User
}

func NewUser(repo repositoty.User) *User {
	return &User{
		repo: repo,
	}
}

func (u *User) Get(user *domain.User) (*domain.User, error) {
	data, err := u.repo.Get(user)
	if err != nil {
		return nil, err
	}
	if err = CheckValidPassword(data.Password, user.Password); err != nil {
		return nil, domain.InvalidPassword
	}
	return u.repo.Get(user)
}

func (u *User) Post(user *domain.User) error {
	user.Id = MakeUuid()
	hashPassword, err := MackeHashPassord(user.Password)
	if err != nil {
		return domain.InternalError
	}
	user.Password = hashPassword
	return u.repo.Post(user)
}

func (u *User) Put(user *domain.User) error {
	hashPassword, err := MackeHashPassord(user.Password)
	if err != nil {
		return domain.InternalError
	}
	user.Password = hashPassword
	return u.repo.Put(user)
}

func (u *User) Delete(user *domain.User) error {
	return u.repo.Delete(user)
}

func MakeUuid() string {
	newuuid := uuid.New()
	return newuuid.String()
}

func MackeHashPassord(value string) (string, error) {
	cost := 10
	hash, err := bcrypt.GenerateFromPassword([]byte(value), cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckValidPassword(oldPassword string, newPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(oldPassword), []byte(newPassword))
	return err
}
