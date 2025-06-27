package application

import "API-VITALVEST/sesion/domain"

type UpdateUser struct {
	repo domain.Isesion
}

func NewUpdateSesion(repo domain.Isesion)*UpdateUser{
	return &UpdateUser{repo: repo}
}


func (c *UpdateUser)Execute(sesion domain.Sesion,id int)error{
	return c.repo.Update(sesion,id)
}