package application

import "API-VITALVEST/sesion/domain"

type GetAllUser struct {
	repo domain.Isesion
}

func NewGetAllUser(repo domain.Isesion) *GetAllUser {
	return &GetAllUser{repo: repo}
}

func (c *GetAllUser) Execute() ([]domain.Sesion, error) {
	return c.repo.GetAll()
}