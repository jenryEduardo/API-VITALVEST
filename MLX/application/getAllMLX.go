package application

import "API-VITALVEST/MLX/domain"

type GetAllMlx_UC struct {
	db domain.IMlx
}

func NewGetAllMlx_UC(db domain.IMlx) *GetAllMlx_UC {
	return &GetAllMlx_UC{db: db}
}

func (uc *GetAllMlx_UC) Run() ([]domain.Mlx, error) {
	return uc.db.FindAll()
}