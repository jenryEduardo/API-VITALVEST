package application

import "API-VITALVEST/BME/domain"

type GetAllBME_UC struct {
	db domain.IBme
}

func NewGetAllBME_UC(db domain.IBme) *GetAllBME_UC {
	return &GetAllBME_UC{db: db}
}

func (uc *GetAllBME_UC) Run() ([]domain.Bme, error) {
	return uc.db.FindAll()
}