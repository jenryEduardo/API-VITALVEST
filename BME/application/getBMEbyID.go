package application

import "API-VITALVEST/BME/domain"

type GetBMEbyID_UC struct {
	db domain.IBme
}

func NewGetBMEbyID(db domain.IBme) *GetBMEbyID_UC {
	return &GetBMEbyID_UC{db: db}
}

func (uc *GetBMEbyID_UC) Run(id int) ([]domain.Bme, error) {
	return uc.db.FindByID(id)
}