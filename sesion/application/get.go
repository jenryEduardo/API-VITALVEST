package application

import "API-VITALVEST/sesion/domain"

type GetUser struct {
	repo domain.Isesion
}

func NewGetUser(repo domain.Isesion)*GetUser{
	return &GetUser{repo: repo}
}

func (c *GetUser)Execute(id int)([]domain.Sesion,error){
	return c.repo.GetById(id)
}