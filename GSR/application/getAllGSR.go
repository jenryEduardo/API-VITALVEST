package application

import "API-VITALVEST/GSR/domain"

type GetAllGsr_UC struct {
	db domain.IGsr
}

func NewGetAllGsr_UC(db domain.IGsr) *GetAllGsr_UC {
	return &GetAllGsr_UC{db: db}
}

func (uc *GetAllGsr_UC) Run() ([]domain.Gsr, error) {
	return uc.db.FindAll()
}
