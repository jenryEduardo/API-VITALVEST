package application

import "API-VITALVEST/MLX/domain"

type DeleteMLX struct {
	db domain.IMlx
}

func NewDeleteMLX(db domain.IMlx) *DeleteMLX {
	return &DeleteMLX{db: db}
}

func (uc *DeleteMLX) Run(id int) error {
	return uc.db.DeleteByID(id)
}