package application

import "API-VITALVEST/sesion/domain"

type SaveUser struct {
	repo domain.Isesion
}

func NewSaveUser(repo domain.Isesion)*SaveUser{
	return &SaveUser{repo: repo}
}

func (c *SaveUser)Execute(sesion domain.Sesion)error{
	return c.repo.Save(sesion)
}