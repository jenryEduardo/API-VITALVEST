package application

import "API-VITALVEST/MPU/domain"

type SaveMPU_UC struct {
	db domain.IMpu
}

func NewSaveMPU_UC (db domain.IMpu) *SaveMPU_UC {
	return &SaveMPU_UC{db: db}
}

func (uc *SaveMPU_UC) Run(MPU domain.Mpu) error {
	return uc.db.Save(MPU)
} 