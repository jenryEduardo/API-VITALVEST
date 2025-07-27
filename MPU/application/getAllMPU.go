package application

import "API-VITALVEST/MPU/domain"

type GetAllMPU_UC struct {
	db domain.IMpu
}

func NewGetAllMPU_UC(db domain.IMpu) *GetAllMPU_UC {
	return &GetAllMPU_UC{db: db}
}

func (uc *GetAllMPU_UC) Run() (int,[]domain.Mpu,error) {
	return uc.db.FindAll()
}
