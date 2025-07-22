package application

import "API-VITALVEST/alertas/domain"

type UpdateAlerta struct {
	db domain.IAlertas
}

func NewUpdateAlerta(db domain.IAlertas) *UpdateAlerta {
	return &UpdateAlerta{db: db}
}

func (uc *UpdateAlerta) Run(id int, Alerta domain.Alerta) error {
	return uc.db.UpdateByID(id, Alerta)
}