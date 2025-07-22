package domain

import (
	"errors"
	"time"
)

type Alerta struct {
	ID                      int       `json:"id"`
	NombreDelSensor         string    `json:"nombre_del_sensor"`
	Fecha                   time.Time `json:"fecha"`
	CantidadDeVecesEnviado  int       `json:"cantidad_de_veces_enviado"`
}

// Valores permitidos para NombreDelSensor
var SensoresPermitidos = map[string]bool{
	"temperatura":            true,
	"hidratacion":            true,
	"temperatura ambiental":  true,
	"pasos":                  true,
}

// Función para validar el sensor
func (a *Alerta) Validar() error {
	if !SensoresPermitidos[a.NombreDelSensor] {
		return errors.New("sensor no válido: " + a.NombreDelSensor)
	}
	return nil
}
