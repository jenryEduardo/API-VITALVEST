package application


import "API-VITALVEST/MPU/domain"

type GetAllTableMPU_UC struct {
	db domain.IMpu
}

func NewGetMPU_UC(db domain.IMpu) *GetAllTableMPU_UC {
	return &GetAllTableMPU_UC{db: db}
}

func (uc *GetAllTableMPU_UC) Run() ([]domain.Mpu, error) {
	return uc.db.GetAll()
}
