package domain

type DataSession struct {
	SessionID         int       `json:"session_id"`
	UserID            int       `json:"user_id"`
	FechaInicio       string `json:"fecha_inicio"`
	FechaFin          string `json:"fecha_fin"`
	TempAmbiente      float32   `json:"temp_ambiente"`
	HumedadRelativa   float32   `json:"humedad_relativa"`
	Conductancia      float32   `json:"conductancia"`
	EstadoHidratacion string    `json:"estado_hidratacion"`
	TempCorporal      float32   `json:"temp_corporal"`
	AceleracionX      float32   `json:"aceleracion_x"`
	AceleracionY      float32   `json:"aceleracion_y"`
	AceleracionZ      float32   `json:"aceleracion_z"`
	Pasos             int       `json:"pasos"`
	NivelActividad    string    `json:"nivel_actividad"`
}
