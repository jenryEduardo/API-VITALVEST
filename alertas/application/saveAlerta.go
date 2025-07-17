package application

import "API-VITALVEST/alertas/domain"

type SaveAlerta_UC struct {
	db domain.IAlertas
}

func NewSaveAlerta_UC (db domain.IAlertas) *SaveAlerta_UC {
	return &SaveAlerta_UC{db: db}
}

func (uc *SaveAlerta_UC) Run(Alerta domain.Alerta) error {
	return uc.db.Save(Alerta)
}
