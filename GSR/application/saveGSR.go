package application

import "API-VITALVEST/GSR/domain"

type SaveGSR_UC struct {
	db domain.IGsr
}

func NewSaveGSR_UC (db domain.IGsr) *SaveGSR_UC {
	return &SaveGSR_UC{db: db}
}

func (uc *SaveGSR_UC) Run(gsr domain.Gsr) error {
	return uc.db.Save(gsr)
} 