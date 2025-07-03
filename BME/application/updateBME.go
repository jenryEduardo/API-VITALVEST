package application

import "API-VITALVEST/BME/domain"

type UpdateBME struct {
	db domain.IBme
}

func NewUpdateBME(db domain.IBme) *UpdateBME {
	return &UpdateBME{db: db}
}

func (uc *UpdateBME) Run(id int, Bme domain.Bme) error {
	return uc.db.UpdateByID(id, Bme)
}