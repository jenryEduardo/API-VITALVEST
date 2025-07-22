package application

import "API-VITALVEST/BME/domain"

type SaveBME_UC struct {
	db domain.IBme
}

func NewSaveBME_UC (db domain.IBme) *SaveBME_UC {
	return &SaveBME_UC{db: db}
}

func (uc *SaveBME_UC) Run(Bme domain.Bme) error {
	return uc.db.Save(Bme)
} 