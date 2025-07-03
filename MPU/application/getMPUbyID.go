package application

import "API-VITALVEST/MPU/domain"

type GetMPUbyID_UC struct {
	db domain.IMpu
}

func NewGetMPUbyID(db domain.IMpu) *GetMPUbyID_UC {
	return &GetMPUbyID_UC{db: db}
}

func (uc *GetMPUbyID_UC) Run(id int) ([]domain.Mpu, error) {
	return uc.db.FindByID(id)
}