package application

import "API-VITALVEST/alertas/domain"

type GetAllAlertas_UC struct {
	db domain.IAlertas
}

func NewGetAllAlertas_UC(db domain.IAlertas) *GetAllAlertas_UC {
	return &GetAllAlertas_UC{db: db}
}

func (uc *GetAllAlertas_UC) Run() ([]domain.Alerta, error) {
	return uc.db.FindAll()
}
