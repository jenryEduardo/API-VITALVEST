package application

import "API-VITALVEST/alertas/domain"

type GetAlertaByID struct {
	db domain.IAlertas
}

func NewGetAlertabyID(db domain.IAlertas) *GetAlertaByID {
	return &GetAlertaByID{db: db}
}

func (uc *GetAlertaByID) Run(id int) ([]domain.Alerta, error) {
	return uc.db.FindByID(id)
}