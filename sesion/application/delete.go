package application

import "API-VITALVEST/sesion/domain"

type Delete struct {
	repo domain.Isesion
}

func NewDelete(repo domain.Isesion)*Delete{
	return &Delete{repo: repo}
}

func (c *Delete)Execute(id int)error{
	return c.repo.Delete(id)
}