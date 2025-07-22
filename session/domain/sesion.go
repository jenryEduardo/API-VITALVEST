package domain

import "time"

type Session struct {
	ID          int       `json:"id" db:"id"`                     // ID de la sesión
	UserID      int       `json:"user_id" db:"id_usuario"`        // ID del usuario que inició la sesión
	FechaInicio time.Time `json:"fecha_inicio" db:"fecha_inicio"` // Inicio de la sesión
	FechasFin    time.Time `json:"fecha_fin" db:"fecha_fin"`       // Fin de la sesión (puedes actualizarlo después)
	Descripcion string    `json:"descripcion" db:"descripcion"`   // Opcional: para dar contexto
}
