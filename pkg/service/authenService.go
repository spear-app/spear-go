package service

import (
	"github.com/spear-app/spear-go/pkg/domain/authen"
	"github.com/spear-app/spear-go/pkg/domain/user"
)

type AuthenService interface {
	Signup(*user.User) error
	Login(*user.User) error
	ReadUserByID(*user.User) error
	Update(*user.User) error
	Delete(string) error
}

type DefaultAuthenService struct {
	repo authen.AuthenRepository
}

func (s DefaultAuthenService) Signup(user *user.User) error {
	return s.repo.Signup(user)
}

func (s DefaultAuthenService) Login(user *user.User) error {
	return s.repo.Login(user)
}

func (s DefaultAuthenService) ReadUserByID(user *user.User) error {
	return s.repo.ReadUserByID(user)
}

func (s DefaultAuthenService) Update(user *user.User) error {
	return s.repo.Update(user)
}

func (s DefaultAuthenService) Delete(id string) error {
	return s.repo.Delete(id)
}
