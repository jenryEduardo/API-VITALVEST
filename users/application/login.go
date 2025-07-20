package application

import "API-VITALVEST/users/domain"

type Login struct {
	repo domain.Iuser
}

func NewLogin(repo domain.Iuser) *Login {
	return &Login{repo: repo}
}

func(c *Login) Execute(name string) (*domain.User, error){
	return c.repo.Login(name)
}