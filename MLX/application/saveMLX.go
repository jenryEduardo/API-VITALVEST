package application

import "API-VITALVEST/MLX/domain"

type SaveMLX_uc struct {
	db domain.IMlx
}

func NewSaveMLX_uc(db domain.IMlx) *SaveMLX_uc {
	return &SaveMLX_uc{db: db}
}

func (uc *SaveMLX_uc) Run(mlx domain.Mlx) error {
	return uc.db.Save(mlx)
}