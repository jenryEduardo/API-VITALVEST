package application

import "API-VITALVEST/MLX/domain"

type UpdateMLX struct {
	db domain.IMlx
}

func NewUpdateMLX(db domain.IMlx) *UpdateMLX {
	return &UpdateMLX{db: db}
}

func (uc *UpdateMLX) Run(id int, mlx domain.Mlx) error {
	return uc.db.UpdateByID(id, mlx)
}