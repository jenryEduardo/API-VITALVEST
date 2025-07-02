package application

import "API-VITALVEST/GSR/domain"

type GetGSRbyID_UC struct {
	db domain.IGsr
}

func NewGetGsrbyID(db domain.IGsr) *GetGSRbyID_UC {
	return &GetGSRbyID_UC{db: db}
}

func (uc *GetGSRbyID_UC) Run(id int) ([]domain.Gsr, error) {
	return uc.db.FindByID(id)
}