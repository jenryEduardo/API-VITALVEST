package application

import "API-VITALVEST/MPU/domain"

type UpdateMPU struct {
	db domain.IMpu
}

func NewUpdateMPU(db domain.IMpu) *UpdateMPU {
	return &UpdateMPU{db: db}
}

func (uc *UpdateMPU) Run(id int, MPU domain.Mpu) error {
	return uc.db.UpdateByID(id, MPU)
}