package application

import "API-VITALVEST/users/domain"

type DeleteUser struct {
	repo domain.Iuser
}

func NewDelete(repo domain.Iuser)*DeleteUser{
	return &DeleteUser{repo: repo}
}

func(c *DeleteUser)Execute(id int)error{
	return c.repo.Delete(id)
}