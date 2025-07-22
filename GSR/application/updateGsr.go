package application

import "API-VITALVEST/GSR/domain"

type UpdateGSR struct {
	db domain.IGsr
}

func NewUpdateGSR(db domain.IGsr) *UpdateGSR {
	return &UpdateGSR{db: db}
}

func (uc *UpdateGSR) Run(id int, gsr domain.Gsr) error {
	return uc.db.UpdateByID(id, gsr)
}