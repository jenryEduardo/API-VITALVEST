package domain

type Mpu struct {
	Id int `json:"id"` 
	Aceleracion_x float64 `json:"aceleracion_x"`
	Aceleracion_y float64 `json:"aceleracion_y"`
	Aceleracion_z float64 `json:"aceleracion_z"`
	Pasos int 			  `json:"pasos"`
	Nivel_actividad string `json:"nivel_actividad"`
}
