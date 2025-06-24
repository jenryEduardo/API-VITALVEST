package application

import "API-VITALVEST/users/domain"

type PostUser struct {
	repo domain.Iuser
}

func NEWUSER(repo domain.Iuser)*PostUser{
	return &PostUser{repo: repo}
}

func(c *PostUser) Execute(user domain.User)error{
	return c.repo.Save(user)
}