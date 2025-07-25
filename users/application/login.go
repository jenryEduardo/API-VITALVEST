package application

import "API-VITALVEST/users/domain"

type Login struct {
	repo domain.Iuser
}

func NewLogin(repo domain.Iuser) *Login {
	return &Login{repo: repo}
}

func(c *Login) Execute(username, password string) (*domain.User, error){
	return c.repo.Login(username, password)
}