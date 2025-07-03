package application

import "API-VITALVEST/MPU/domain"

type DeleteMPU struct {
	db domain.IMpu
}

func NewDeleteMPU(db domain.IMpu) *DeleteMPU {
	return &DeleteMPU{db: db}
}

func (uc *DeleteMPU) Run(id int) error {
	return uc.db.DeleteByID(id)
}