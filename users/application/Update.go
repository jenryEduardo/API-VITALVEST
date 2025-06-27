package application

import "API-VITALVEST/users/domain"

type UpdateUser struct {
	repo domain.Iuser
}

func NewUpdate(repo domain.Iuser)*UpdateUser{
	return &UpdateUser{repo: repo}
}

func (c *UpdateUser)Execute(user domain.User,id int)error{
	return c.repo.Update(user,id)
}