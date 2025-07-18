package application

import (
	"API-VITALVEST/alertas/domain"
	"log"
)

type SaveAlerta_UC struct {
	db domain.IAlertas
	notifier domain.Notifier
}

func NewSaveAlerta_UC (db domain.IAlertas, notifier domain.Notifier) *SaveAlerta_UC {
	return &SaveAlerta_UC{db: db, notifier: notifier}
}

func (uc *SaveAlerta_UC) Run(Alerta domain.Alerta) error {
	return uc.db.Save(Alerta)
	if err := uc.notifier.EnviarAlerta(Alerta); err != nil{
		log.Printf("Error notificando alerta: %v", err)
	}
	return nil
}
