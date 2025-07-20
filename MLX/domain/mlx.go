package domain

type Mlx struct {
	Id int `json:"id"` 
	TemperaturaAmbiente float64 `json:"temperatura_ambiente"`
	TemperaturaObjeto float64 `json:"temperatura_objeto"`
}
