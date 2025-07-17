package application

import "API-VITALVEST/alertas/domain"

type DeleteAlerta struct {
	db domain.IAlertas
}

func NewDeleteAlerta(db domain.IAlertas) *DeleteAlerta {
	return &DeleteAlerta{db: db}
}

func (uc *DeleteAlerta) Run(id int) error {
	return uc.db.DeleteByID(id)
}