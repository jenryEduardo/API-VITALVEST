package application

import "API-VITALVEST/BME/domain"

type DeleteBME struct {
	db domain.IBme
}

func NewDeleteBME(db domain.IBme) *DeleteBME {
	return &DeleteBME{db: db}
}

func (uc *DeleteBME) Run(id int) error {
	return uc.db.DeleteByID(id)
}