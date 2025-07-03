package application

import "API-VITALVEST/MLX/domain"

type GetMlxbyID_UC struct {
	db domain.IMlx
}

func NewGetMlxbyID(db domain.IMlx) *GetMlxbyID_UC {
	return &GetMlxbyID_UC{db: db}
}

func (uc *GetMlxbyID_UC) Run(id int) ([]domain.Mlx, error) {
	return uc.db.FindByID(id)
}