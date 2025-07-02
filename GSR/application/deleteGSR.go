package application

import "API-VITALVEST/GSR/domain"

type DeleteGSR struct {
	db domain.IGsr
}

func NewDeleteGSR(db domain.IGsr) *DeleteGSR {
	return &DeleteGSR{db: db}
}

func (uc *DeleteGSR) Run(id int) error {
	return uc.db.DeleteByID(id)
}